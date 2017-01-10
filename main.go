package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// TODO use this later when event schema is clear
type issueEvent struct {
	action string `json:action`
	issue  struct {
		title  string `json:title`
		labels []struct {
			name string `json:name`
		} `json:labels`
		assignee   string `json:assignee`
		created_at string `json:created_at`
		updated_at string `json:updated_at`
		body       string `json:body`
	} `json:issue`
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", eventHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func eventHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}
