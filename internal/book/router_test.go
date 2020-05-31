// TODO. write api test
// GET /api/books
// POST /api/books/{id}
// PATCH /api/books/{id}
// DELETE /api/books/{id}

package book

import (
	"log"
	"net/http"
	"net/http/httptest"
	"sync"
)

var (
	serverAddr string
	once       sync.Once
)

func startServer() {
	mux := http.NewServeMux()

	Router(mux)

	server := httptest.NewServer(mux)
	serverAddr = server.Listener.Addr().String()
	log.Print("Test WebSocket server listening on ", serverAddr)
}
