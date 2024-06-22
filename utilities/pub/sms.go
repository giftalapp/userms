package pub

import (
	"crypto/rand"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/giftalapp/authsrv/utilities/bucket"
	"github.com/pquerna/otp/totp"
)

type SMS struct {
	PubService
	sc     *sns.Client
	bucket *bucket.Bucket
}

func (s *SMS) Send(phoneNumber string) (string, error) {
	log.Printf("[PUB] Send SMS to %s", phoneNumber)

	secret := make([]byte, 64)

	_, err := rand.Read(secret)
	if err != nil {
		return "", err
	}

	code, err := totp.GenerateCode(string(secret), time.Now())

	if err != nil {
		return "", err
	}

	log.Printf("[PUB] SEND OTP: %s\n", code)

	s.bucket.Set(phoneNumber, string(secret))
	//TODO: generate a encrypted token from the phone number

	return phoneNumber, nil
}
