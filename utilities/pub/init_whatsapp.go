package pub

import (
	"net/http"

	"github.com/giftalapp/userms/config"
	"github.com/piusalfred/whatsapp"
)

func initWhatsApp() (*whatsapp.Client, error) {
	client := whatsapp.NewClient(
		whatsapp.WithHTTPClient(http.DefaultClient),
		whatsapp.WithPhoneNumberID(config.Env.WhatsAppPhoneNumberID),
		whatsapp.WithBusinessAccountID(config.Env.WhatsAppID),
		whatsapp.WithAccessToken(config.Env.WhatsAppToken),
	)

	return client, nil
}
