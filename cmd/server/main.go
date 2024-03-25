package main

import (
	"fmt"
)

func run() error {
	fmt.Println("starting app")
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println("error: ", err)
	}
}
