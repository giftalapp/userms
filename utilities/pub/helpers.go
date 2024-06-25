package pub

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
	"math"
	"strconv"
	"time"

	"github.com/giftalapp/userms/config"
	"github.com/giftalapp/userms/utilities/bucket"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pquerna/otp/hotp"
)

func getSHA256(data string) (string, error) {
	h := sha256.New()
	io.Copy(
		h,
		bytes.NewReader([]byte(data)),
	)
	hSum := base64.URLEncoding.EncodeToString(h.Sum(nil))

	return hSum, nil
}

func createOtpAndToken(buck *bucket.Bucket, phoneNumber string) (string, string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":   config.Env.AppName,
		"phone": phoneNumber,
	})

	signedToken, err := token.SignedString([]byte(config.Env.JWTSecret))

	if err != nil {
		return "", "", errors.New("server_jwt_sign_error")
	}

	signedTokenHash, err := getSHA256(signedToken)

	if err != nil {
		return "", "", errors.New("server_hash_error")
	}

	if _, _, err := buck.Get(signedTokenHash); err == nil {
		return "", "", errors.New("already_in_verification")
	}

	key, err := hotp.Generate(hotp.GenerateOpts{
		Issuer:      config.Env.AppName,
		AccountName: phoneNumber,
	})

	if err != nil {
		return "", "", errors.New("server_otp_generation_error")
	}

	otp, err := hotp.GenerateCode(
		key.Secret(),
		1,
	)

	if err != nil {
		return "", "", errors.New("server_otp_generation_error")
	}

	if err = buck.Set(signedTokenHash, key.Secret(), time.Duration(0)); err != nil {
		return "", "", errors.New("server_bucket_set_error")
	}

	if err = buck.Set("counter_"+signedTokenHash, 1, time.Duration(0)); err != nil {
		return "", "", errors.New("server_bucket_set_error")
	}

	return otp, signedToken, err
}

func updateOtpCounter(buck *bucket.Bucket, signedToken string) (string, error) {
	signedTokenHash, err := getSHA256(signedToken)

	if err != nil {
		return "", errors.New("server_hash_error")
	}

	var otpSecret interface{}
	var ttl time.Duration
	if otpSecret, ttl, err = buck.Get(signedTokenHash); err != nil {
		return "", errors.New("token_expired_error")
	}

	var otpCounter interface{}
	if otpCounter, _, err = buck.Get("counter_" + signedTokenHash); err != nil {
		return "", errors.New("token_expired_error")
	}

	counter, _ := strconv.ParseUint(otpCounter.(string), 10, 64)

	passedSeconds := uint64((config.Env.OTPExpire - ttl).Seconds())
	refreshIn := uint64(config.Env.OTPRefresh.Seconds()) * counter

	if refreshIn > passedSeconds {
		return "", errors.New("otp_not_expired_error")
	}

	counter += uint64(math.Floor(float64(passedSeconds)/float64(refreshIn)) + 1)

	if err := buck.Set("counter_"+signedTokenHash, counter, ttl); err != nil {
		return "", errors.New("bucket_set_error")
	}

	otp, err := hotp.GenerateCode(
		otpSecret.(string),
		counter,
	)

	if err != nil {
		return "", errors.New("server_otp_generation_error")
	}

	return otp, err
}
