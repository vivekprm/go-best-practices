package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	err := doThis()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("handling %q: %v", r.RequestURI, err)
		return
	}

	err = doThat()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("handling %q: %v", r.RequestURI, err)
		return
	}
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
