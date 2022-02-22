package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type interview struct {
	Message   string
	Timestamp int
}

func main() {
	request()
}

func youtubeinterview(response http.ResponseWriter, r *http.Request) {
	interviewtype := []interview{
		interview{
			Message:   "Hello from Michael Haidner",
			Timestamp: 01152020,
		},
	}

	json.NewEncoder(response).Encode(interviewtype)

	fmt.Println("Endpoint", interviewtype)
}

func healthcheck(response http.ResponseWriter, r *http.Request) {

	response.WriteHeader(http.StatusOK)

	return
}

func request() {

	//http.HandleFunc("/", healthcheck)

	http.HandleFunc("/interview", youtubeinterview)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
