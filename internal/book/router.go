package book

import (
	"io"
	"net/http"
)

// Router default
func Router() {
	http.HandleFunc("/book", func(responseWriter http.ResponseWriter, request *http.Request) {
		io.WriteString(responseWriter, "Hello Book from a HandleFunc #1!\n")
	})
}
