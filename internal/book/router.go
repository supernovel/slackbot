package book

import (
	"io"
	"net/http"
)

// Router default
func Router(mux *http.ServeMux) {
	mux.HandleFunc("/api/books", func(responseWriter http.ResponseWriter, request *http.Request) {
		io.WriteString(responseWriter, "Hello Book from a HandleFunc!\n")
		io.WriteString(responseWriter, request.URL.Path)
		if request.Method != http.MethodPost {
			// TODO: Redirect 405
		}
		// TODO: Get book list from database
	})
	// TODO: /api/books/{id} --> No id --> Redirect to /api/books
}
