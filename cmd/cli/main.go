package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	"github.com/headshed-dev/queue-lite/internal/db"
	"github.com/headshed-dev/queue-lite/internal/queue"
)

func run() error {
	fmt.Println("starting consumer app")

	// Read the scripts file
	data, err := ioutil.ReadFile("scripts.json")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Unmarshal the JSON data
	var scripts map[string]string
	if err := json.Unmarshal(data, &scripts); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	db, error := db.NewDatabase()
	if error != nil {
		fmt.Printf("failed to create database connection:\n")
		return error
	}
	if err := db.Ping(context.Background()); err != nil {
		fmt.Printf("failed to ping database:\n")
		return err
	}
	queueService := queue.NewService(db)
	newJob, err := queueService.ConsumeJobs(context.Background())
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return err
	}
	fmt.Printf("job: [%s]\n", newJob.Payload)

	// scripts := map[string]string{
	// 	"script2":                   "/path/to/script2",
	// 	"script3":                   "/path/to/script3",
	// 	"PostExportJob-cms-lite001": "/path/to/script1",
	// }

	script, ok := scripts[string(newJob.Payload)]
	if !ok {
		fmt.Printf("script not found: %s\n", newJob.Payload)
		return errors.New("script not found for payload : " + string(newJob.Payload))
	}
	log.Printf("found script: %s in lookup\n", script)

	// Print the value of 'script'
	fmt.Printf("Value of 'script': %s\n", script)

	// Run the script
	cmd := exec.Command(script)
	output, err := cmd.CombinedOutput()

	// Check the exit code
	exitCode := cmd.ProcessState.ExitCode()
	if exitCode == 0 {
		// Script exited normally
		fmt.Printf("Script exited with code 0\n")
		fmt.Printf("Output: %s\n", output)
	} else {
		// Script exited with an error
		fmt.Printf("Script exited with code %d\n", exitCode)
		log.Fatalf("Script error: %s", output)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println("error: ", err)
	}
}
