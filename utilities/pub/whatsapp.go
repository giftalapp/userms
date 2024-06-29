package pub

import (
	"context"
	"fmt"

	"github.com/giftalapp/userms/utilities/bucket"
	"github.com/piusalfred/whatsapp"
	"github.com/piusalfred/whatsapp/models"
)

type WhatsApp struct {
	PubService
	wa     *whatsapp.Client
	bucket *bucket.Bucket
}

func (w *WhatsApp) sendWhatsAppOtp(phoneNumber string, otp string, lang string) error {
	_, err := w.wa.SendTextTemplate(context.Background(), phoneNumber, &whatsapp.TextTemplateRequest{
		Name:         "en_verification_code",
		LanguageCode: lang,
		Body: []*models.TemplateParameter{
			{
				Type: "text",
				Text: otp,
			},
		},
	})

	if err != nil {
		err = fmt.Errorf("server_error %s", err)
	}

	return err
}

func (w *WhatsApp) Send(phoneNumber string, lang string) (string, error) {
	otp, token, err := createOtpAndToken(w.bucket, phoneNumber)

	if err != nil {
		return "", err
	}

	err = w.sendWhatsAppOtp(phoneNumber, otp, lang)

	return token, err
}

func (w *WhatsApp) Resend(phoneNumber string, lang string) error {
	otp, err := updateOtpCounter(w.bucket, phoneNumber)

	if err != nil {
		return err
	}

	err = w.sendWhatsAppOtp(phoneNumber, otp, lang)

	return err
}
