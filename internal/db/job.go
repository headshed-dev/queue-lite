package db

import (
	"context"
	"log"
)

type Job struct {
	Name    string
	Payload string
}

func (d *Database) AddJob(
	ctx context.Context,
	name string, payload []byte,
) (uint64, error) {

	jobId, e := d.Client.Put(0, 0, 60, []byte(payload))
	if e != nil {
		log.Fatal(e)
	}

	return jobId, nil
}
