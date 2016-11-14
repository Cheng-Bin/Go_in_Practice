package chapter1

import (
	"io/ioutil"
	"net/http"
)

func HttpGet() string {
	resp, _ := http.Get("http://www.baidu.com")
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return string(body)
}
