package src

import "net/http"

type RouteHandler struct {
}

func NewRouteHandler() *RouteHandler {
	return &RouteHandler{}
}

func (rh *RouteHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /login", rh.handleLogin)
}

func (rh *RouteHandler) handleLogin(w http.ResponseWriter, r *http.Request) {

}
