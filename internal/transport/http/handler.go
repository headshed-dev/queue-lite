package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

// type JobService interface{}

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
	h.Router.HandleFunc("/alive", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("running"))
	}).Methods("GET")

	h.Router.HandleFunc("/api/v1/job", h.PostJob).Methods("POST")
	// h.Router.HandleFunc("/api/v1/job", h.GetJobs).Methods("GET")
}

func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println("failed to start server", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)
	log.Println("shutting downn service")
	return nil
}
