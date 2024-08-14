package main

import (
	"fmt"
	"net/http"
	"time"
)

// func greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World! %s", time.Now())
// }

func main() {
	http.HandleFunc("/", helloMyName)
	http.ListenAndServe(":8080", nil)
}

func helloMyName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", time.Now())
}
