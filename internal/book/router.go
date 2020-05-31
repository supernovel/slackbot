package book

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/slack-go/slack"
)

// Router : apply router for slack
func Router(mux *http.ServeMux) {
	mux.HandleFunc("/api/books", func(responseWriter http.ResponseWriter, request *http.Request) {
		log.Printf("Path => %s Method => %s", request.URL.Path, request.Method)

		if request.Method != http.MethodPost {
			responseWriter.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		s, err := slack.SlashCommandParse(request)
		if err != nil {
			responseWriter.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !s.ValidateToken(os.Getenv("SLACK_VERIFICATION_TOKEN")) {
			responseWriter.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Header Section
		headerText := slack.NewTextBlockObject("mrkdwn", "도서 목록입니다.", false, false)
		headerSection := slack.NewSectionBlock(headerText, nil, nil)

		// Fields
		typeField := slack.NewTextBlockObject("mrkdwn", "*러스트 프로그래밍 공식 가이드*\n:star::star::star::star: 1528 reviews\n 스티브 클라브닉 , 캐롤 니콜스 지음 | 장현희 옮김 | 제이펍 | 2019년 11월 28일 출간\n`대여가능`", false, false)
		bookImage := slack.NewImageBlockElement("http://image.kyobobook.co.kr/images/book/large/729/l9791188621729.jpg", "Windsor Court Hotel thumbnail")
		fieldsSection := slack.NewSectionBlock(typeField, nil, slack.NewAccessory(bookImage))

		// Approve and Deny Buttons
		previousBtnTxt := slack.NewTextBlockObject("plain_text", "Prev", false, false)
		previousBtn := slack.NewButtonBlockElement("", "click_me_123", previousBtnTxt)

		nextBtnTxt := slack.NewTextBlockObject("plain_text", "Next", false, false)
		nextBtn := slack.NewButtonBlockElement("", "click_me_123", nextBtnTxt)

		actionBlock := slack.NewActionBlock("", previousBtn, nextBtn)

		// Build Message with blocks created above

		msg := slack.NewBlockMessage(
			headerSection,
			fieldsSection,
			actionBlock,
		)

		responseWriter.Header().Set("Content-Type", "application/json")
		json.NewEncoder(responseWriter).Encode(msg)
		// TODO: get book list from database
	})

	// TODO: /api/books/ : add, update, delete book from database
}
