package users

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	handler *Handler
}

func NewRouter(h *Handler) http.Handler {
	router := &Router{
		handler: h,
	}

	return router.InitRoutes()
}

func (r *Router) InitRoutes() http.Handler{
	m := mux.NewRouter()

	m.Handle("/registration", r.handler.SignUp()).Methods(http.MethodPost)
	m.Handle("/authorization", r.handler.SignIn()).Methods(http.MethodPost)

	return m
}