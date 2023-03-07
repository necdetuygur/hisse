package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

const userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36"

func Handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	q := req.URL.Query().Get("q")
	fmt.Fprintf(w, GetPrice(q))
}

func GetRequest(url string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", userAgent)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func GetPrice(q string) string {
	str := GetRequest("https://finance.yahoo.com/quote/" + q + ".IS/")
	str = Parse(str, `active="">(.*?)<\/fin-streamer>`, 1, "")
	return str
}

func Parse(str string, rgx string, key int, clr string) string {
	r, _ := regexp.Compile(rgx)
	arr := r.FindStringSubmatch(str)
	if len(arr) == 0 {
		return "0"
	}
	pri := strings.ReplaceAll(strings.TrimSpace(arr[key]), clr, "")
	return pri
}
