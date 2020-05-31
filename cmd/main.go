package main

import (
	"io"
	"log"
	"net/http"

	"github.com/supernovel/slackbot/internal/book"
)

func defaultHandler(responseWriter http.ResponseWriter, request *http.Request) {
	io.WriteString(responseWriter, "404 Not Found!\n")
}

func main() {
	addr := ":8080"

	http.HandleFunc("/", defaultHandler)
	book.Router()

	log.Printf("Listening... => %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
