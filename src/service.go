package src

import (
	"database/sql"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"github.com/giftalapp/userms/utilities/pub"
)

type AuthService struct {
	addr string
	db   *sql.DB
	fb   *firebase.App
	pubc *pub.Pub
}

func NewAuthService(addr string, db *sql.DB, fb *firebase.App, pubc *pub.Pub) *AuthService {
	return &AuthService{
		addr: addr,
		db:   db,
		fb:   fb,
		pubc: pubc,
	}
}

func (srv *AuthService) Run() error {
	handler := http.NewServeMux()

	routeHandler := NewRouteHandler(srv.db, srv.fb, srv.pubc)
	routedHandler, err := routeHandler.RegisterRoutes(handler)

	if err != nil {
		return err
	}

	server := &http.Server{
		Addr:    srv.addr,
		Handler: routedHandler,
	}

	return server.ListenAndServe()
}
