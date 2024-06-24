package middleware

import (
	"context"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/appcheck"
)

type AppLimitter struct {
	handler http.Handler
	check   *appcheck.Client
}

func NewAppLimitter(handler http.Handler, fb *firebase.App) (*AppLimitter, error) {
	check, err := fb.AppCheck(context.Background())

	if err != nil {
		return nil, err
	}

	return &AppLimitter{
		handler: handler,
		check:   check,
	}, nil
}

func (al *AppLimitter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if appCheckToken, ok := r.Header[http.CanonicalHeaderKey("X-Firebase-AppCheck")]; ok {
		_, err := al.check.VerifyToken(appCheckToken[0])

		if err != nil && appCheckToken[0] != "debugKey" { // TODO: REMOVE DEBUG KEY OVERRIDE
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized.\n"))
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized.\n"))
		return
	}

	al.handler.ServeHTTP(w, r)
}
