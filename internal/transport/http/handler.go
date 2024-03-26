package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

type JobService interface{}

type Handler struct {
	Router  *mux.Router
	Service JobService
	Server  *http.Server
}

func NewHandler(service JobService) *Handler {
	h := &Handler{
		Service: service,
	}

	h.Router = mux.NewRouter()

	h.mapRoutes()
	h.Server = &http.Server{
		Addr:    ":8080",
		Handler: h.Router,
	}
	return h
}

func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	}).Methods("GET")
}

func (h *Handler) Serve() error {
	if err := h.Server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
