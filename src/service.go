package src

import (
	"database/sql"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type AuthService struct {
	addr string
	db   *sql.DB
	sns  *sns.Client
}

func NewAuthService(addr string, db *sql.DB, sns *sns.Client) *AuthService {
	return &AuthService{
		addr: addr,
		db:   db,
		sns:  sns,
	}
}

func (srv *AuthService) Run() error {
	handler := http.NewServeMux()

	server := &http.Server{
		Addr:    srv.addr,
		Handler: handler,
	}

	routeHandler, err := NewRouteHandler(srv.sns)
	if err != nil {
		return err
	}

	routeHandler.RegisterRoutes(handler)

	return server.ListenAndServe()
}
