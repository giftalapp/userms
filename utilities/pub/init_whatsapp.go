package pub

import (
	"github.com/giftalapp/userms/config"
	"github.com/husseinamine/whatsapp"
)

func initWhatsApp() (*whatsapp.Client, error) {
	client := whatsapp.NewClient(config.Env.WhatsAppPhoneNumberID, config.Env.WhatsAppToken)

	return client, nil
}
