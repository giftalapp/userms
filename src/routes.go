package src

import (
	"database/sql"
	"net/http"

	"github.com/giftalapp/authsrv/src/handlers/verification"
	"github.com/giftalapp/authsrv/src/middleware"
	"github.com/giftalapp/authsrv/utilities/pub"
)

type RouteHandler struct {
	db   *sql.DB
	pubc *pub.Pub
}

func NewRouteHandler(db *sql.DB, pubc *pub.Pub) *RouteHandler {
	return &RouteHandler{
		db:   db,
		pubc: pubc,
	}
}

func (rh *RouteHandler) RegisterRoutes(router *http.ServeMux) http.Handler {
	router.HandleFunc("POST /verification/send", verification.SendHandler)
	router.HandleFunc("POST /verification/resend", verification.ResendHandler)
	router.HandleFunc("POST /verification/verify", verification.VerifyHandler)

	injectedHandler := middleware.NewServiceInjector(router, rh.db, rh.pubc)
	rateLimitedHandler := middleware.NewRateLimitter(injectedHandler)

	return rateLimitedHandler
}
