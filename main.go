package main

import (
	"io"
	"net/http"
	"os"
	"strings"
)

var res = strings.Trim(`
{"name": "gorilla", "age": 27, "likes": ["dog", "cat", "gorilla"]}
`, "\n")

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, res)
	})

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
