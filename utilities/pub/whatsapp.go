package pub

import (
	"log"

	"github.com/giftalapp/authsrv/utilities/bucket"
)

type WhatsApp struct {
	PubService
	bucket *bucket.Bucket
}

func (s *WhatsApp) Send(phoneNumber string) (string, error) {
	otp, token, err := createOtpAndToken(s.bucket, phoneNumber)

	if err != nil {
		return "", err
	}

	log.Printf("[PUB] Send WhatsApp to %s", phoneNumber)
	log.Printf("[PUB] SENT OTP: %s\n", otp)

	return token, nil
}
