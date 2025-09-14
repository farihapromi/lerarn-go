package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "i am promi", "I am a SW Engineer")
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)

}
