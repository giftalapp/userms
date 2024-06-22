package src

import (
	"database/sql"
	"net/http"

	"github.com/giftalapp/authsrv/utilities/pub"
)

type AuthService struct {
	addr string
	db   *sql.DB
	pubc *pub.Pub
}

func NewAuthService(addr string, db *sql.DB, pubc *pub.Pub) *AuthService {
	return &AuthService{
		addr: addr,
		db:   db,
		pubc: pubc,
	}
}

func (srv *AuthService) Run() error {
	handler := http.NewServeMux()

	routeHandler := NewRouteHandler(srv.db, srv.pubc)
	routedHandler := routeHandler.RegisterRoutes(handler)

	server := &http.Server{
		Addr:    srv.addr,
		Handler: routedHandler,
	}

	return server.ListenAndServe()
}
