package main

import (
	"fmt"
	"handle"
	"net/http"
	"strconv"
	"vo"
)

//default_status:503:Service Unavailable
var monitor = vo.Monitor{http.StatusServiceUnavailable, "503 Service Unavailable"}

func main() {
	http.HandleFunc("/", HttpStatusHandler)
	http.HandleFunc("/deploy", handle.DeployHandler)
	var healthCheck handle.HealthCheck
	healthCheck.Monitor = monitor
	http.Handle("/hc", healthCheck)

	http.ListenAndServe("localhost:8000", nil)
}

func HttpStatusHandler(writer http.ResponseWriter, request *http.Request) {
	code, err := strconv.Atoi(request.URL.Query().Get("code"))
	if err == nil {
		monitor.Status = code
		writer.WriteHeader(monitor.Status)
		fmt.Fprintln(writer, "Hello, Http Status:", monitor.Status)
	} else {
		writer.WriteHeader(monitor.Status)
		fmt.Fprintln(writer, "Hello, Http Status:", monitor.Status)
	}

}
