package db

/*
type Job struct {
	Name    string `json:"name"`
	Payload string `json:"payload"`
}

// AddJob adds a job to the database
func (d *Database) AddJob(
	ctx context.Context,
	name string, payload []byte,
	tubeName string,
) (uint64, error) {

	if tubeName == "" {
		tubeName = "default"
	}

	log.Println("Adding job to tube:", tubeName)

	d.Client.Use(tubeName)
	jobId, e := d.Client.Put(0, 0, 60, []byte(payload))
	if e != nil {
		log.Fatal(e)
	}

	return jobId, nil
}

*/
