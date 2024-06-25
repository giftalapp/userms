package middleware

import (
	"context"
	"net/http"

	"github.com/giftalapp/userms/utilities/pub"
	"github.com/jackc/pgx/v5"
)

type ServiceKey int

const (
	DBKey  ServiceKey = 0
	PubKey ServiceKey = 1
)

func GetDB(r *http.Request) *pgx.Conn {
	return r.Context().Value(DBKey).(*pgx.Conn)
}

func GetPub(r *http.Request) *pub.Pub {
	return r.Context().Value(PubKey).(*pub.Pub)
}

type ServiceInjector struct {
	handler http.Handler
	db      *pgx.Conn
	pubc    *pub.Pub
}

func NewServiceInjector(handler http.Handler, db *pgx.Conn, pubc *pub.Pub) *ServiceInjector {
	return &ServiceInjector{
		handler: handler,
		db:      db,
		pubc:    pubc,
	}
}

func (i *ServiceInjector) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), DBKey, i.db)
	ctx = context.WithValue(ctx, PubKey, i.pubc)

	i.handler.ServeHTTP(w, r.WithContext(ctx))
}
