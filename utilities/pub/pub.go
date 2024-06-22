package pub

import (
	"time"

	"github.com/giftalapp/authsrv/utilities/bucket"
)

type PubService interface {
	Send(string) (string, error)
}

type Pub struct {
	SMS      *SMS
	WhatsApp *WhatsApp
}

func NewPubClient(redisURL string) (*Pub, error) {
	sc, err := initSNS()

	if err != nil {
		return nil, err
	}

	bucket, err := bucket.NewBucket(redisURL, time.Second*30)

	if err != nil {
		return nil, err
	}

	return &Pub{
		SMS: &SMS{
			sc:     sc,
			bucket: bucket,
		},
		WhatsApp: &WhatsApp{},
	}, nil
}
