package src

import (
	"database/sql"
	"net/http"
)

type AuthService struct {
	addr string
	db   *sql.DB
}

func NewAuthService(addr string, db *sql.DB) *AuthService {
	return &AuthService{
		addr: addr,
		db:   db,
	}
}

func (srv *AuthService) Run() error {
	handler := http.NewServeMux()

	server := &http.Server{
		Addr:    srv.addr,
		Handler: handler,
	}

	routeHandler := NewRouteHandler()
	routeHandler.RegisterRoutes(handler)

	return server.ListenAndServe()
}
