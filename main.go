package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", MonitorHandler)
	http.ListenAndServe("localhost:8000", nil)
}

func MonitorHandler(writer http.ResponseWriter, request *http.Request) {
	request.pa
	writer.WriteHeader(400)
	fmt.Fprintln(writer, "Hello , Mars")
}
