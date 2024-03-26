package queue

import (
	"context"
	"errors"
	"fmt"
	"log"
)

var (
	ErrSubmittingPayload = errors.New("failed to submit queue payload")
	ErrNotImplemented    = errors.New("not implemented")
)

// Job is a struct that represents a job payload
type Job struct {
	Name    string
	Payload []byte
}

// Service is a struct that represents the Queue Service
type Service struct {
	Store Store
}

// Store - defines method for storing queue payloads
type Store interface {
	PostPayload(ctx context.Context, job Job) (Job, error)
	ConsumeJobs(ctx context.Context) (Job, error)
}

// NewQueue is a constructor for the Queue Service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) PostJob(ctx context.Context, job Job) (Job, error) {

	log.Printf("posting job: [%s]\n", job.Payload)

	insertedQueue, err := s.Store.PostPayload(ctx, job)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return Job{}, ErrSubmittingPayload
	}
	return insertedQueue, nil

}

func (s *Service) ConsumeJobs(ctx context.Context) (Job, error) {

	log.Printf("consuming jobs\n")
	readJob, err := s.Store.ConsumeJobs(ctx)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return Job{}, ErrNotImplemented
	}
	return readJob, nil

}

func (s *Service) ListJobs(ctx context.Context) ([]Job, error) {
	return []Job{}, ErrNotImplemented
}
