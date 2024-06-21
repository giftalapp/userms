package src

import (
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/giftalapp/authsrv/src/handlers"
)

type RouteHandler struct {
	sns *sns.Client
}

func NewRouteHandler(sns *sns.Client) (*RouteHandler, error) {
	return &RouteHandler{
		sns: sns,
	}, nil
}

// Custom auth-to-handler compatibility
func (rh *RouteHandler) asAuthHandler(authHandler func(http.ResponseWriter, *http.Request, *sns.Client)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authHandler(w, r, rh.sns)
	}
}

func (rh *RouteHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /verify", rh.asAuthHandler(handlers.VerifyTokenHandler))
}
