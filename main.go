package main

import (
	"fmt"
	"net/http"
)	

func main(){
	port := "8080"
	fmt.Println("Starting app")
	HttpServerExample(port)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "{\"message\":\"hello\"}")
}

func HttpServerExample(port string) {
	const host string = ":"

	http.HandleFunc("/hello", hello)

	fmt.Println("server starting at", port)
	http.ListenAndServe(host + port, nil)
}