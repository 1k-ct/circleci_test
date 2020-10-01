package main

import (
	"fmt"
	"net/http"
)

func hoge(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hoge no mi")
}
func top(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "my top page hello")
}
func main() {
	http.HandleFunc("/", top)
	http.HandleFunc("/hoge", hoge)
	http.ListenAndServe(":8080", nil)
}
