package handle

import (
	"fmt"
	"net/http"
	"vo"
)

type HealthCheck struct {
	vo.Monitor
}

func (h HealthCheck) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	_, err := http.Get("http://localhost:8080")
	if err == nil {
		h.Status = http.StatusOK
		h.Desc = "OK"
	} else {
		h.Status = http.StatusServiceUnavailable
		h.Desc = "Service Unavailable"
	}
	fmt.Fprintln(writer, "OK")
}
