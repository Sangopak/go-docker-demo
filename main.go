package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := "8080"
	fmt.Println("Starting app")
	HttpServerExample(port)
}

func Hello() string {
	return "{\"message\":\"hello\"}"
}

func Okay() string {
	return "ok"
}

func healthCheckHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, Okay())
}
func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, Hello())
}

func HttpServerExample(port string) {
	const host string = ":"

	http.HandleFunc("/", healthCheckHandler)
	http.HandleFunc("/healthcheck", healthCheckHandler)
	http.HandleFunc("/api/hello", helloHandler)

	fmt.Println("server starting at", port)
	http.ListenAndServe(host+port, nil)
}
