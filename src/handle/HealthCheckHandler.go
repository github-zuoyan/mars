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
	resp, err := http.Get("http://0.0.0.0:8080")
	if err == nil && resp != nil && resp.StatusCode == http.StatusOK {
		h.Monitor.Status = http.StatusOK
		h.Monitor.Desc = "OK"
	} else {
		h.Monitor.Status = http.StatusServiceUnavailable
		h.Monitor.Desc = "Service Unavailable"
	}
	if resp != nil {
		resp.Body.Close()
	}
	fmt.Fprintln(writer, "http_status:", h.Monitor.Status)
}
