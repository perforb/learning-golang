package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintf(w, "hello\n")
	if err != nil {
		panic(err)
	}
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, header := range req.Header {
		for _, h := range header {
			_, err := fmt.Fprintf(w, "%v: %v\n", name, h)
			if err != nil {
				panic(err)
			}
		}
	}
}

// $ curl localhost:8090/headers                                                                 (git)-[master]
// User-Agent: curl/7.78.0
// Accept: */*
func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}
