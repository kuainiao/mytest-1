package main

import (
	"net/http"
	"fmt"
	"html"
)

func main() {
	http.HandleFunc("/", Test)
	http.ListenAndServe(":8888", nil)

}

func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world!")
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
