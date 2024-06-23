package fb

import (
	"context"

	firebase "firebase.google.com/go/v4"
)

func Start() (*firebase.App, error) {
	return firebase.NewApp(context.Background(), nil)
}
