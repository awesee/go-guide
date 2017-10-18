package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	httpGet("http://www.baidu.com")
}

func httpGet(url string) {
	resp, err := http.Get(url)
	check(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
