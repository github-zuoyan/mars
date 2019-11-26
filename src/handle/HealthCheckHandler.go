package handle

import (
	"fmt"
	"net/http"
	"vo"
)

type HealthCheck struct {
	Monitor *vo.Monitor
}

func (h HealthCheck) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	_, err := http.Get("http://localhost:8080")
	if err == nil {
		h.Monitor.Status = http.StatusOK
		h.Monitor.Desc = "OK"
	} else {
		h.Monitor.Status = http.StatusServiceUnavailable
		h.Monitor.Desc = "Service Unavailable"
	}
	fmt.Fprintln(writer, "http_status:", h.Monitor.Status)
}
