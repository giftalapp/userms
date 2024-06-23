package pub

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"time"

	"github.com/giftalapp/authsrv/config"
	"github.com/giftalapp/authsrv/utilities/bucket"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pquerna/otp/totp"
)

type PubService interface {
	Send(string) (string, error)
	Resend(string)
}

type Pub struct {
	SMS      *SMS
	WhatsApp *WhatsApp
}

func NewPubClient(redisURL string) (*Pub, error) {
	sc, err := initSNS()

	if err != nil {
		return nil, err
	}

	bucket, err := bucket.NewBucket(redisURL, config.Env.RedisExpire)

	if err != nil {
		return nil, err
	}

	return &Pub{
		SMS: &SMS{
			sc:     sc,
			bucket: bucket,
		},
		WhatsApp: &WhatsApp{
			bucket: bucket,
		},
	}, nil
}

func getSHA256(data string) (string, error) {
	h := sha256.New()
	io.Copy(
		h,
		bytes.NewReader([]byte(data)),
	)
	hSum := base64.URLEncoding.EncodeToString(h.Sum(nil))

	return hSum, nil
}

func createOtpAndToken(bucket *bucket.Bucket, phoneNumber string) (string, string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":   config.Env.AppName,
		"phone": phoneNumber,
	})

	signedToken, err := token.SignedString([]byte(config.Env.JWTSecret))

	if err != nil {
		return "", "", err
	}

	signedTokenHash, err := getSHA256(signedToken)

	if err != nil {
		return "", "", err
	}

	if err := bucket.Get(signedTokenHash); err == nil {
		return "", "", fmt.Errorf("phone number is already being verified")
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      config.Env.AppName,
		AccountName: phoneNumber,
	})

	if err != nil {
		return "", "", err
	}

	otp, err := totp.GenerateCode(key.Secret(), time.Now())

	if err != nil {
		return "", "", err
	}

	err = bucket.Set(signedTokenHash, key.Secret())

	return otp, signedToken, err
}
