package pub

import (
	"log"

	"github.com/giftalapp/userms/utilities/bucket"
)

type WhatsApp struct {
	PubService
	bucket *bucket.Bucket
}

func (w *WhatsApp) Send(phoneNumber string) (string, error) {
	otp, token, err := createOtpAndToken(w.bucket, phoneNumber)

	if err != nil {
		return "", err
	}

	log.Printf("[PUB] Send WhatsApp to %s", phoneNumber)
	log.Printf("[PUB] SENT OTP: %s\n", otp)

	return token, nil
}

func (w *WhatsApp) Resend(phoneNumber string) error {
	otp, err := updateOtpCounter(w.bucket, phoneNumber)

	if err != nil {
		return err
	}

	log.Printf("[PUB] Re-Send SMS to %s", phoneNumber)
	log.Printf("[PUB] SENT OTP: %s\n", otp)

	return nil
}
