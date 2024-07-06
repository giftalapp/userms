package pub

import (
	"errors"
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

	res, err := w.wa.SendMessage(message.Message)

	if _, ok := res["error"]; ok {
		err = fmt.Errorf("server_error %s", res["error"])
	}

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

	// Remove the verification token from cache if OTP send fails
	if err != nil {
		signedTokenHash, err := getSHA256(token)

		if err != nil {
			return "", errors.New("server_hash_error")
		}

		w.bucket.Del(signedTokenHash)
		w.bucket.Del("counter_" + signedTokenHash)
	}

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
