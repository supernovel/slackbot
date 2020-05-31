package book

import (
	"io"
	"net/http"
)

// Router default
func Router(mux *http.ServeMux) {
	mux.HandleFunc("/book/", func(responseWriter http.ResponseWriter, request *http.Request) {
		io.WriteString(responseWriter, "Hello Book from a HandleFunc!\n")
		io.WriteString(responseWriter, request.URL.Path)
	})
}
