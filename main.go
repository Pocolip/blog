package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

// This is a method named save that takes as its receiver p, a pointer to Page.
// It takes no parameters, and returns a value of type error.
func (b *Blog) save() error {
	filename := b.Title + ".json"
	return ioutil.WriteFile(filename, []byte(b.Content), 0600)
}

func main() {
	router := NewRouter()
	log.Println("Listening on port 80 (Standard HTTP)...")
	log.Fatal(http.ListenAndServe(":80", router))
}
