package handle

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func HeartHandler(writer http.ResponseWriter, request *http.Request) {
	values := request.URL.Query()
	c, err := strconv.Atoi(values.Get("c"))
	if err != nil {
		c = 10
	}

	n, err := strconv.Atoi(values.Get("n"))
	if err != nil {
		n = 1000
	}

	url := values.Get("url")

	if url == "" {
		url = "http://0.0.0.0:8080"
	}

	ch := make(chan int, n)

	count := n / c

	var wg sync.WaitGroup
	wg.Add(c + 1)
	for i := 0; i < c; i++ {
		go GetHttp(url, count, ch, &wg)
	}

	go GetHttp(url, n-count*c, ch, &wg)

	wg.Wait()
	close(ch)
	sum := 0
	for v := range ch {
		sum += v
	}

	fmt.Fprint(writer, fmt.Sprintf("time=%.2f count=%d qps=%.2f", float64(sum)/1000000.0, n, float64(n)*1000000.0/float64(sum)))

}

func GetHttp(url string, n int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		start := time.Now().UnixNano()
		http.Get(url)
		end := time.Now().UnixNano()
		ch <- int((end - start) / 1000)
		if end <= start {
			println("result=", (end-start)/1000)
		}
	}
}
