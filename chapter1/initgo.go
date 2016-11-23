package chapter1

import (
	"fmt"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello, My name is InitGo Montoya.")
}

func Server() {
	http.HandleFunc("/", hello)
	http.ListenAndServe("localhost:4000", nil)
}
