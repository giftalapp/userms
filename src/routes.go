package src

import (
	"database/sql"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"github.com/giftalapp/authsrv/src/handlers/verification"
	"github.com/giftalapp/authsrv/src/middleware"
	"github.com/giftalapp/authsrv/utilities/pub"
)

type RouteHandler struct {
	db   *sql.DB
	fb   *firebase.App
	pubc *pub.Pub
}

func NewRouteHandler(db *sql.DB, fb *firebase.App, pubc *pub.Pub) *RouteHandler {
	return &RouteHandler{
		db:   db,
		fb:   fb,
		pubc: pubc,
	}
}

func (rh *RouteHandler) RegisterRoutes(router *http.ServeMux) (http.Handler, error) {
	router.HandleFunc("POST /verification/send", verification.SendHandler)
	router.HandleFunc("POST /verification/resend", verification.ResendHandler)
	router.HandleFunc("POST /verification/verify", verification.VerifyHandler)

	injectedHandler := middleware.NewServiceInjector(router, rh.db, rh.pubc)
	appLimitedHandler, err := middleware.NewAppLimitter(injectedHandler, rh.fb)

	if err != nil {
		return nil, err
	}

	rateLimitedHandler := middleware.NewRateLimitter(appLimitedHandler)

	return rateLimitedHandler, nil
}
