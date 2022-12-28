package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/", errorHandler(betterHandler))
}

func errorHandler(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("handling %q: %v", r.RequestURI, err)
		}
	}
}

func betterHandler(w http.ResponseWriter, r *http.Request) error {
	if err := doThis(); err != nil {
		return fmt.Errorf("doing this: %v", err)
	}

	if err := doThat(); err != nil {
		return fmt.Errorf("doing that: %v", err)
	}
	return nil
}

func main() {
	fmt.Printf("Starting server at port 8000\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

func doThis() error {
	fmt.Println("Doing this")
	return nil
}

func doThat() error {
	fmt.Println("Doing that")
	return errors.New("error in doing that")
}
