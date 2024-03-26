package http

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/headshed-dev/queue-lite/internal/queue"
)

type JobJSON struct {
	Name    string `json:"Name"`
	Payload string `json:"Payload"`
}

type JobService interface {
	PostJob(context.Context, queue.Job) (queue.Job, error)
}

func (h *Handler) PostJob(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read request body:", err)
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var job JobJSON
	err = json.Unmarshal(body, &job)
	if err != nil {
		log.Println("Failed to decode JSON body:", err)
		http.Error(w, "Failed to decode JSON body", http.StatusBadRequest)
		return
	}

	convertedJob := queue.Job{
		Name:    job.Name,
		Payload: []byte(job.Payload),
	}

	submittedJob, err := h.Service.PostJob(r.Context(), convertedJob)
	if err != nil {
		log.Println("Failed to post job:", err)
		http.Error(w, "Failed to post job", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(submittedJob); err != nil {
		log.Println("Failed to encode job:", err)
		http.Error(w, "Failed to encode job", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetJobs(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}
