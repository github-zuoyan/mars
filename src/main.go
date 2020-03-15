package main

import (
	"fmt"
	"handle"
	"net/http"
	"strconv"
	"time"
	"vo"
)

//default_status:503:Service Unavailable
var monitor = vo.Monitor{http.StatusServiceUnavailable, "503 Service Unavailable"}

func main() {

	//启动定时任务,检测后端应用如果异常，则自动设置状态为503
	go func() {
		ticker := time.NewTicker(time.Millisecond * 500)
		for {
			select {
			case <-ticker.C:
				resp, err := http.Get("http://0.0.0.0:8080")
				if err != nil || resp != nil && resp.StatusCode != http.StatusOK {
					monitor.Status = http.StatusServiceUnavailable
					monitor.Desc = "Service Unavailable"
				}
				if resp != nil {
					resp.Body.Close()
				}
			}
		}
	}()
	http.HandleFunc("/", HttpStatusHandler)
	http.HandleFunc("/deploy", handle.DeployHandler)
	http.HandleFunc("/heat", handle.HeatHandler)
	var healthCheck handle.HealthCheck
	healthCheck.Monitor = &monitor
	http.Handle("/hc", healthCheck)

	server := http.Server{Addr: "0.0.0.0:8000", Handler: nil, ReadTimeout: time.Second * 3}
	server.SetKeepAlivesEnabled(false)
	server.ListenAndServe()

}

func HttpStatusHandler(writer http.ResponseWriter, request *http.Request) {
	code, err := strconv.Atoi(request.URL.Query().Get("code"))
	if err == nil {
		monitor.Status = code
		writer.WriteHeader(monitor.Status)
		fmt.Fprintln(writer, "Http_Status:", monitor.Status)
	} else {
		writer.WriteHeader(monitor.Status)
		fmt.Fprintln(writer, "Http_Status:", monitor.Status)
	}

}
