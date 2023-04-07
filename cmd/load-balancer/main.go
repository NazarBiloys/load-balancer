package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello")
	log.Infof("OPA")
	if err != nil {
		log.Error(err)
	}
	return
}

func main() {
	log.Infof("Start application..")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	handler := http.HandlerFunc(handler)
	http.Handle("/", handler)

	log.Printf("Starting server on port %s\n", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Errorf(err.Error())
	}
}
