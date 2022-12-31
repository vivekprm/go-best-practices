package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	jobs := []string{"one", "two", "three"}

	errc := make(chan error)
	for _, job := range jobs {
		// Calling concurrently
		go func(job string) {
			errc <- do(job)
		}(job)
	}
	for range jobs {
		if err := <-errc; err != nil {
			fmt.Println(err)
		}
	}
}

func do(job string) error {
	fmt.Println("doing job", job)
	time.Sleep(1 * time.Second)
	return errors.New("something went wrong")
}
