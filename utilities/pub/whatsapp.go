package pub

import "log"

type WhatsApp struct {
	PubService
}

func (s *WhatsApp) Send(phoneNumber string) (string, error) {
	log.Printf("[PUB] Send WhatsApp to %s", phoneNumber)

	return "", nil
}
