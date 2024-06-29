package src

import (
	"net/http"

	firebase "firebase.google.com/go/v4"
	"github.com/giftalapp/userms/src/handlers/verification"
	"github.com/giftalapp/userms/src/middleware"
	"github.com/giftalapp/userms/utilities/pub"
	"github.com/jackc/pgx/v5"
)

type RouteHandler struct {
	db   *pgx.Conn
	fb   *firebase.App
	pubc *pub.Pub
}

func NewRouteHandler(db *pgx.Conn, fb *firebase.App, pubc *pub.Pub) *RouteHandler {
	return &RouteHandler{
		db:   db,
		fb:   fb,
		pubc: pubc,
	}
}

func (rh *RouteHandler) RegisterRoutes(router *http.ServeMux) (http.Handler, error) {
	router.HandleFunc("POST /api/user/verification/send", verification.SendHandler)
	router.HandleFunc("POST /api/user/verification/resend", verification.ResendHandler)
	router.HandleFunc("POST /api/user/verification/verify", verification.VerifyHandler)

	injectedHandler := middleware.NewServiceInjector(router, rh.db, rh.pubc)
	appLimitedHandler, err := middleware.NewAppLimitter(injectedHandler, rh.fb)

	if err != nil {
		return nil, err
	}

	rateLimitedHandler := middleware.NewRateLimitter(appLimitedHandler)

	return rateLimitedHandler, nil
}
