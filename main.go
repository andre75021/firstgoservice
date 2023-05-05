package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

func showInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Current time", time.Now())
	fmt.Fprintln(w, "URL path: ", html.EscapeString(r.URL.Path))
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	http.HandleFunc("/", showInfo)
	http.HandleFunc("/site", serveFile)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe ", err)
	}
}
