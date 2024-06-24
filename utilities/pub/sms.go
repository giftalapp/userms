package pub

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/giftalapp/authsrv/utilities/bucket"
)

type SMS struct {
	PubService
	sc     *sns.Client
	bucket *bucket.Bucket
}

func (s *SMS) Send(phoneNumber string) (string, error) {
	otp, token, err := createOtpAndToken(s.bucket, phoneNumber)

	if err != nil {
		return "", err
	}

	log.Printf("[PUB] Send SMS to %s", phoneNumber)
	log.Printf("[PUB] SENT OTP: %s\n", otp)

	return token, nil
}

func (s *SMS) Resend(phoneNumber string) error {
	otp, err := updateOtpCounter(s.bucket, phoneNumber)

	if err != nil {
		return err
	}

	log.Printf("[PUB] Re-Send SMS to %s", phoneNumber)
	log.Printf("[PUB] SENT OTP: %s\n", otp)

	return nil
}
