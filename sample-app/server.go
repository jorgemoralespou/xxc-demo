package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
	log.Print("Request received")
	remote_service, exist := os.LookupEnv("REMOTE_SERVICE")
	if exist {
		log.Print("Calling remote service as env REMOTE_SERVICE has been defined")
		resp, err := http.Get("http://" + remote_service + "/hello")
		if err != nil {
			fmt.Fprintf(w, "Empty answer from remote service. Probably not running or bad configuration")
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		fmt.Fprintf(w, "Answer from remote service: %q", body)
	} else {
		log.Print("Local answer as env REMOTE_SERVICE has NOT been defined")
		fmt.Fprintf(w, "hello")
	}

}

func main() {
	http.HandleFunc("/", hello)
	port, exist := os.LookupEnv("PORT")
	if exist {
		log.Print("Server started")
		http.ListenAndServe(":"+port, nil)
	} else {
		log.Print("No PORT env defined")
	}
}
