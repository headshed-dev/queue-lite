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

// Service is the main struct that represents the Queue Service and is returned by the NewService constructor
type Service struct {
	Store Store
}

// Store - defines method for storing queue payloads and is an interface passed to NewService
type Store interface {
	PostPayload(ctx context.Context, job Job) (Job, error)
	ConsumeJobs(ctx context.Context) (Job, error)
}

// NewQueue is a constructor for the Queue Service, it accepts a store interface and returns a new Queue Service struct
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

// PostJob is a method on the Queue Service that accepts a job payload and returns a Job struct
func (s *Service) PostJob(ctx context.Context, job Job) (Job, error) {

	log.Printf("posting job: [%s]\n", job.Payload)

	insertedQueue, err := s.Store.PostPayload(ctx, job)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return Job{}, ErrSubmittingPayload
	}
	return insertedQueue, nil

}

// ConsumeJobs is a method on the Queue Service that returns a Job struct
func (s *Service) ConsumeJobs(ctx context.Context) (Job, error) {

	log.Printf("consuming jobs\n")
	readJob, err := s.Store.ConsumeJobs(ctx)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return Job{}, ErrNotImplemented
	}
	return readJob, nil

}

// ListJobs is a method on the Queue Service that returns a list of Job structs
func (s *Service) ListJobs(ctx context.Context) ([]Job, error) {
	return []Job{}, ErrNotImplemented
}
