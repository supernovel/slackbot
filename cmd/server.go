package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/supernovel/slackbot/internal/book"
)

func defaultHandler(responseWriter http.ResponseWriter, request *http.Request) {
	io.WriteString(responseWriter, "404 Not Found!\n")
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	addr := ":" + port
	mux := http.NewServeMux()

	mux.HandleFunc("/", defaultHandler)
	book.Router(mux)

	log.Printf("Listening... => %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
