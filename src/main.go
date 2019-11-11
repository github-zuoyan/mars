package main

import (
	"fmt"
	"handle"
	"net/http"
	"strconv"
)

type Monitor struct {
	status int
	desc   string
}

//default_status:503:Service Unavailable
var monitor = Monitor{http.StatusServiceUnavailable, "503 Service Unavailable"}

func main() {
	http.HandleFunc("/", HttpStatusHandler)
	http.HandleFunc("/deploy", handle.DeployHandler)

	http.ListenAndServe("localhost:8000", nil)
}

func HttpStatusHandler(writer http.ResponseWriter, request *http.Request) {
	code, err := strconv.Atoi(request.URL.Query().Get("code"))
	if err == nil {
		monitor.status = code
		writer.WriteHeader(monitor.status)
		fmt.Fprintln(writer, "Hello, Http Status:", monitor.status)
	} else {
		writer.WriteHeader(monitor.status)
		fmt.Fprintln(writer, "Hello, Http Status:", monitor.status)
	}

}
