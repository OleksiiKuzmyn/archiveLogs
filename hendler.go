package main

import (
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func ArchiveLogsHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	fmt.Println("Start.")

	logs := GetLogs()

	serviceClien := ConnectToStorageAccount("DefaultEndpointsProtocol=https;AccountName=logsol;AccountKey=/4CgsqFT3wX4k3QRLIFsPg086JVPF47NLKQK3mYVxR+X/K98SmH6VJHNnWTl/XFNe4Ty+y4r9jxg+AStVeozBQ==;EndpointSuffix=core.windows.net")

	container := CreateContainerIfNotExists("testcontainer4", serviceClien)

	UploadLogs(container, logs)

	// log.Printf("Hello From - Timer Trigger.\n")

	// data := Response{Message: "Hello World From - Timer Trigger."}
	// json, _ := json.Marshal(data)
	// fmt.Fprint(rw, string(json))

	fmt.Println("End")
}
