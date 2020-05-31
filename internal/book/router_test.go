package book

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"testing"
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

// TestBooks_NotAllowMethod : /api/books
//
// only allow POST method, other method return 405(method not allow)
func TestBooks_NotAllowMethod(t *testing.T) {
	once.Do(startServer)
	serverURL := fmt.Sprintf("http://%s/api/books", serverAddr)

	tests := []struct {
		method         string
		wantStatusCode int
	}{
		{
			method:         http.MethodHead,
			wantStatusCode: http.StatusMethodNotAllowed,
		},
		{
			method:         http.MethodGet,
			wantStatusCode: http.StatusMethodNotAllowed,
		},
		{
			method:         http.MethodPost,
			wantStatusCode: http.StatusOK,
		},
		{
			method:         http.MethodPut,
			wantStatusCode: http.StatusMethodNotAllowed,
		},
		{
			method:         http.MethodDelete,
			wantStatusCode: http.StatusMethodNotAllowed,
		},
		{
			method:         http.MethodTrace,
			wantStatusCode: http.StatusMethodNotAllowed,
		},
	}

	client := &http.Client{}

	for i, test := range tests {
		t.Logf("Method => %s, WantedStatus => %d", test.method, test.wantStatusCode)

		req, err := http.NewRequest(test.method, serverURL, nil)

		if err != nil {
			t.Fatalf("%d: Unexpected error: %s", i, err)
		}

		resp, err := client.Do(req)

		if err != nil {
			t.Fatalf("%d: Unexpected error: %s", i, err)
		}

		if resp.StatusCode != test.wantStatusCode {
			t.Errorf("%d: Got status code %d, want %d", i, resp.StatusCode, test.wantStatusCode)
		}

		resp.Body.Close()
	}
}

// TestBooks_SlackCommand : /api/books
func TestBooks_CheckSlackCommand(t *testing.T) {
	once.Do(startServer)
	serverURL := fmt.Sprintf("http://%s/api/books", serverAddr)

	tests := []struct {
		body           url.Values
		wantStatusCode int
	}{
		{
			body: url.Values{
				"command": []string{"/book"},
				"text":    []string{"list"},
			},
			wantStatusCode: http.StatusOK,
		},
		{
			body: url.Values{
				"command": []string{"/book"},
				"text":    []string{"list kafka"},
			},
			wantStatusCode: http.StatusOK,
		},
		{
			body: url.Values{
				"command": []string{"/book"},
				"text":    []string{"describe kafka"},
			},
			wantStatusCode: http.StatusOK,
		},
		{
			body: url.Values{
				"command": []string{"/book"},
				"text":    []string{"describe"},
			},
			wantStatusCode: http.StatusBadRequest,
		},
		{
			body: url.Values{
				"command": []string{"/book"},
				"text":    []string{"lists"},
			},
			wantStatusCode: http.StatusBadRequest,
		},
		{
			body: url.Values{
				"command": []string{"/books"},
				"text":    []string{"list"},
			},
			wantStatusCode: http.StatusBadRequest,
		},
	}

	client := &http.Client{}

	for i, test := range tests {
		req, err := http.NewRequest(http.MethodPost, serverURL, strings.NewReader(test.body.Encode()))

		if err != nil {
			t.Fatalf("%d: Unexpected error: %s", i, err)
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		resp, err := client.Do(req)

		if err != nil {
			t.Fatalf("%d: Unexpected error: %s", i, err)
		}

		if resp.StatusCode != test.wantStatusCode {
			t.Errorf("%d: Got status code %d, want %d", i, resp.StatusCode, test.wantStatusCode)
		}

		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			t.Fatalf("%d: Unexpected error: %s", i, err)
		}

		t.Logf("Body => %s", body)
	}
}
