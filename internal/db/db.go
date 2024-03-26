package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/headshed-dev/queue-lite/internal/queue"
	"github.com/nutrun/lentil"
)

type Database struct {
	Client *lentil.Beanstalkd
}

func NewDatabase() (*Database, error) {
	// "beanstalkd:11300"
	connectionString := fmt.Sprintf(
		"%s:%s",
		os.Getenv("BEANSTALKD_HOST"),
		os.Getenv("BEANSTALKD_PORT"),
	)

	conn, e := lentil.Dial(connectionString)
	if e != nil {
		return &Database{}, fmt.Errorf("failed to connect to beanstalkd : %w", e)
	}

	return &Database{
		Client: conn,
	}, nil

}

func (d *Database) Ping(ctx context.Context) error {
	_, e := d.Client.Stats()
	if e != nil {
		return fmt.Errorf("failed to ping beanstalkd : %w", e)
	}
	return nil
}

func (db *Database) PostPayload(ctx context.Context, job queue.Job) (queue.Job, error) {
	_, e := db.Client.Put(0, 0, 120, job.Payload)
	if e != nil {
		return queue.Job{}, fmt.Errorf("failed to post payload to beanstalkd : %w", e)
	}
	return job, nil
}

func (db *Database) ConsumeJobs(ctx context.Context) (queue.Job, error) {

	job, e := db.Client.Reserve()
	if e != nil {
		return queue.Job{}, fmt.Errorf("failed to consume job from beanstalkd : %w", e)
	}
	log.Printf("JOB ID: %d, JOB BODY: %s", job.Id, job.Body)
	e = db.Client.Delete(job.Id)
	if e != nil {
		return queue.Job{}, fmt.Errorf("failed to delete job from beanstalkd : %w", e)
	}

	convertedJob := queue.Job{
		Name:    fmt.Sprintf("%d", job.Id),
		Payload: job.Body,
	}

	return convertedJob, nil
}
