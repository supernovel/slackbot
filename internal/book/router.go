package book

import (
	"log"
	"net/http"
)

// Router : apply router for slack
func Router(mux *http.ServeMux) {
	mux.HandleFunc("/api/books", func(responseWriter http.ResponseWriter, request *http.Request) {
		log.Printf("Path => %s Method => %s", request.URL.Path, request.Method)

		if request.Method != http.MethodGet {
			responseWriter.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// TODO: get book list from database
	})

	// TODO: /api/books/ : add, update, delete book from database
}
