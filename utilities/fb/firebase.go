package fb

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func Start() (*firebase.App, error) {
	return firebase.NewApp(context.Background(), nil, option.WithCredentialsFile(".firebase.json"))
}
