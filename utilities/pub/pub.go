package pub

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/giftalapp/userms/config"
	"github.com/giftalapp/userms/utilities/bucket"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pquerna/otp/hotp"
)

type PubService interface {
	Send(string) (string, error)
	Resend(string) error
}

type Pub struct {
	bucket   *bucket.Bucket
	SMS      *SMS
	WhatsApp *WhatsApp
}

func NewPubClient(redisURL string) (*Pub, error) {
	sc, err := initSNS()

	if err != nil {
		return nil, err
	}

	bucket, err := bucket.NewBucket(redisURL, config.Env.OTPExpire)

	if err != nil {
		return nil, err
	}

	return &Pub{
		bucket: bucket,
		SMS: &SMS{
			sc:     sc,
			bucket: bucket,
		},
		WhatsApp: &WhatsApp{
			bucket: bucket,
		},
	}, nil
}

func (p *Pub) Verify(otp string, signedToken string) (string, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("server_invalid_secret_error")
		}

		return []byte(config.Env.JWTSecret), nil
	})

	if err != nil {
		if strings.HasSuffix(err.Error(), "server_") {
			return "", err
		}

		return "", errors.New("invalid_token_error")
	}

	claims := token.Claims.(jwt.MapClaims)

	signedTokenHash, err := getSHA256(signedToken)

	if err != nil {
		return "", fmt.Errorf("server_hash_error %s", err)
	}

	var otpSecret interface{}
	var ttl time.Duration

	if otpSecret, ttl, err = p.bucket.Get(signedTokenHash); err != nil {
		return "", errors.New("token_expired_error")
	}

	otpCounter, _, _ := p.bucket.Get("counter_" + signedTokenHash)
	counter, _ := strconv.ParseUint(otpCounter.(string), 10, 64)

	passedSeconds := uint64((config.Env.OTPExpire - ttl).Seconds())
	refreshIn := uint64(config.Env.OTPRefresh.Seconds()) * counter

	if passedSeconds > refreshIn {
		return "", errors.New("otp_expired_error")
	}

	ok := hotp.Validate(otp, counter, otpSecret.(string))

	if !ok {
		return "", errors.New("invalid_otp_error")
	}

	p.bucket.Del(signedTokenHash)
	p.bucket.Del("counter_" + signedTokenHash)

	return claims["phone"].(string), nil
}
