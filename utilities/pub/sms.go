package pub

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type SMS struct {
	PubService
	sc *sns.Client
}

func (s *SMS) Send(phoneNumber string) (string, error) {
	log.Printf("[PUB] Send SMS to %s", phoneNumber)

	return "", nil
}
