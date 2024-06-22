package pub

type PubService interface {
	Send(string) (string, error)
}

type Pub struct {
	SMS      *SMS
	WhatsApp *WhatsApp
}

func NewPubClient() (*Pub, error) {
	sc, err := initSNS()

	if err != nil {
		return nil, err
	}

	return &Pub{
		SMS: &SMS{
			sc: sc,
		},
		WhatsApp: &WhatsApp{},
	}, nil
}
