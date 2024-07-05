package pub

import (
	"fmt"

	"github.com/giftalapp/userms/utilities/bucket"
	"github.com/husseinamine/whatsapp"
	"github.com/husseinamine/whatsapp/types"
)

type WhatsApp struct {
	PubService
	wa     *whatsapp.Client
	bucket *bucket.Bucket
}

func (w *WhatsApp) sendWhatsAppOtp(phoneNumber string, otp string, lang string) error {
	factory := whatsapp.NewMessageFactory()
	factory.WithTo(phoneNumber)

	message := factory.WithTemplate("verification_code", lang)
	message.WithBody()
	message.WithBodyParamater(&types.Parameter{
		Type: types.TextParameter,
		Text: otp,
	})
	message.WithButton(types.URL)
	message.WithButtonParameter(&types.ButtonParameter{
		Type: types.TextButtonParameter,
		Text: otp,
	})
	message.Save()

	_, err := w.wa.SendMessage(message.Message)

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
