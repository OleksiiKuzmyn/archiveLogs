package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	http.HandleFunc("/", ArchiveLogsHandler)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
