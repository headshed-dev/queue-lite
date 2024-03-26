package db

import (
	"context"
	"fmt"
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

/*
func (d *Database) PostPayload(ctx context.Context, payload []byte) error {
	_, e := d.Client.Put(0, 0, 120, payload)
	if e != nil {
		return fmt.Errorf("failed to post payload to beanstalkd : %w", e)
	}
	return nil
}
*/
