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
		doConcurrently(job, errc)
	}
	for range jobs {
		if err := <-errc; err != nil {
			fmt.Println(err)
		}
	}
}

func doConcurrently(job string, err chan error) {
	go func() {
		fmt.Println("doing job", job)
		time.Sleep(1 * time.Second)
		err <- errors.New("something went wrong!")
	}()
}
