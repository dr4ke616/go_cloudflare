package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func make_request() {

	apiUrl := "https://www.cloudflare.com"
	resource := "/api_json.html"

	data := url.Values{}
	data.Add("email", "accounts@glassrobotstudios.com")
	data.Add("tkn", "d609ccba709d89a59e29be77504796e383503")
	data.Add("a", "zone_load_multi")
	// data.Add("z", "glassrobotstudios.com")

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := fmt.Sprintf("%v", u)

	log.Println(urlStr)
	log.Println(bytes.NewBufferString(data.Encode()))

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
	r.Header.Add("Content-type", "application/x-www-form-urlencoded")
	r.Header.Add("Accept", "text/plain")

	resp, _ := client.Do(r)
	defer resp.Body.Close()

	log.Println(resp.Status)
	contents, _ := ioutil.ReadAll(resp.Body)

	log.Println(string(contents))

}

func main() {
	make_request()
}
