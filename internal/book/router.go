package book

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	slack "github.com/slack-go/slack"
	book "github.com/supernovel/slackbot/internal/book/models"
)

// Router : apply router for slack
func Router(mux *http.ServeMux) {
	mux.HandleFunc("/api/books", func(responseWriter http.ResponseWriter, request *http.Request) {
		log.Printf("/api/books : Path => %s", request.URL.Path)

		if request.Method != http.MethodPost {
			responseWriter.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		slackCommand, err := slack.SlashCommandParse(request)

		if err != nil {
			responseWriter.WriteHeader(http.StatusInternalServerError)
			return
		}

		if slackCommand.Command != "/book" {
			responseWriter.WriteHeader(http.StatusBadRequest)
			return
		}

		// TODO: Remove logging body
		serializedBody, err := json.Marshal(slackCommand)

		if err != nil {
			responseWriter.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Printf("SlackCommand => %s", serializedBody)

		splitedText := strings.Split(slackCommand.Text, " ")

		if splitedText[0] == "list" {
			// TODO: Build book list from database
			message, err := BuildBookListBlock(&[]book.Book{
				{
					ISBN:   "123124124123",
					Title:  "kafka",
					Author: "human",
				},
				{
					ISBN:   "123124124123",
					Title:  "prometheus",
					Author: "robot",
				},
			}, 0, 2)

			if err != nil {
				responseWriter.WriteHeader(http.StatusInternalServerError)
				return
			}

			responseWriter.Header().Set("Content-Type", "application/json")
			responseWriter.Write(message)

			return
		}

		if splitedText[0] == "describe" && len(splitedText) == 2 {
			// TODO: Get book info from database
			return
		}

		responseWriter.WriteHeader(http.StatusBadRequest)
	})
}
