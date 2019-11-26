package handle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	URL "net/url"
	"sync"
	"time"
)

func HeatHandler(writer http.ResponseWriter, request *http.Request) {
	values := request.URL.Query()

	var c, n, url = 10, 1000, "http://0.0.0.0:8080"
	appname := values.Get("appname")

	if appname != "" {
		resp, err := http.Get("http://tool.vip.qiyi.domain/preheat/query?appname=" + appname)

		defer resp.Body.Close()

		if err == nil && resp.StatusCode == 200 {

			body, _ := ioutil.ReadAll(resp.Body)
			var urlMapList []map[string]string
			json.Unmarshal(body, &urlMapList)
			urlMap := urlMapList[0]
			conditionMap := make(map[string]int)
			json.Unmarshal([]byte(urlMap["condition"]), &conditionMap)

			c = conditionMap["c"]

			n = conditionMap["n"]

			url = urlMap["url"]

			if url == "" {
				url = "http://0.0.0.0:8080"
			} else {
				u, err := URL.Parse(url)
				if err != nil {
					url = "http://0.0.0.0:8080/"
				} else {
					url = "http://0.0.0.0:8080" + u.Path + "?" + u.RawQuery
				}
			}
		}
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
	}
}
