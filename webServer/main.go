package main

import (
	"fmt"
	"net/http"
)

func main() {
	startServer()
}

func startServer() {

	http.HandleFunc("/", invokingLambda)

	http.ListenAndServe(":8080", nil)

}

func invokingLambda(_ http.ResponseWriter, r *http.Request) {

	fmt.Printf("Received request with headers: %v\n", r.Header)
	fmt.Printf("Received request with path: %v\n", r.URL.Path)

	for k, v := range r.Header {
		fmt.Printf("Header: %s, Value: %s\n", k, v)
	}
}
