package main

import (
	"fmt"
	"html"
	"net/http"
	"os"

	"github.com/go-kit/log"
)

func main() {
	w := log.NewSyncWriter(os.Stderr)
	logger := log.NewLogfmtLogger(w)
	logger.Log("question", "what is the meaning of life?", "answer", 42)

	// comment
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi There Yolanda!")
	})

	http.ListenAndServe(":8080", nil)
}
