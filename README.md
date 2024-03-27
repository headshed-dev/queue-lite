# queue-lite

simple api to make pushing jobs to a queue via an http rest api

a cli to read jobs ( as consumer ) can be used to process jobs from the same api

# prerequisites 

beanstalkd must be running, in the following example it is on localhost port 11300

golang must be installed and available to run for example the following if you wish to extend the application and woud like to make changes. To see it running, you can use something like :

```
BEANSTALKD_HOST="localhost" BEANSTALKD_PORT="11300" go run cmd/server/main.go
```

run the command line consumer with something lik

```
BEANSTALKD_HOST="localhost" BEANSTALKD_PORT="11300"  go run cmd/cli/main.go 
```

# build

```bash
docker build -t queue-lite:001 .
```

# run 

```bash
docker run -d -p 8080:8080 -e BEANSTALKD_HOST=localhost -e BEANSTALKD_PORT=11300 queue-lite:001
```

# Tests

integration tests run with

```bash
BEANSTALKD_HOST="localhost" BEANSTALKD_PORT="11300" go test -tags=integration -v ./...
```

