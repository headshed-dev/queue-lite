package main

import (
	"context"
	"fmt"

	"github.com/headshed-dev/queue-lite/internal/db"
	"github.com/headshed-dev/queue-lite/internal/queue"
	transportHttp "github.com/headshed-dev/queue-lite/internal/transport/http"
)

func run() error {
	fmt.Println("starting app")

	db, error := db.NewDatabase()
	if error != nil {
		fmt.Printf("failed to create database connection:\n")
		return error
	}
	if err := db.Ping(context.Background()); err != nil {
		fmt.Printf("failed to ping database:\n")
		return err
	}

	fmt.Println("app started with database connection")
	queueService := queue.NewService(db)
	// queueService.PublishPayload(context.Background(), queue.Job{
	// 	Name:    "test",
	// 	Payload: []byte("test payload 534532453"),
	// })

	httpHandler := transportHttp.NewHandler(queueService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println("error: ", err)
	}
}
