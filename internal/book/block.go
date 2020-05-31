package book

import (
	"encoding/json"

	"github.com/slack-go/slack"
	book "github.com/supernovel/slackbot/internal/book/models"
)

// BuildBookListBlock is build interactive book list block
func BuildBookListBlock(list *[]book.Book, current int, total int) ([]byte, error) {
	var bookListBlocks []slack.Block

	// Header block
	headerText := slack.NewTextBlockObject("mrkdwn", "Found Book 4 / 30 (2)", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	bookListBlocks = append(bookListBlocks, headerSection)
	bookListBlocks = append(bookListBlocks, slack.NewDividerBlock())

	// Body block
	for _, book := range *list {
		bookTextBlocks := make([]*slack.TextBlockObject, 0)
		bookTextBlocks = append(bookTextBlocks, slack.NewTextBlockObject("mrkdwn", "*<fakeLink.com|"+book.Title+">*", false, false))
		bookTextBlocks = append(bookTextBlocks, slack.NewTextBlockObject("mrkdwn", "*"+book.Author+"*", false, false))

		bookListBlocks = append(bookListBlocks, slack.NewSectionBlock(nil, bookTextBlocks, nil))
		bookListBlocks = append(bookListBlocks, slack.NewDividerBlock())
	}

	// Action block
	nextBtnTxt := slack.NewTextBlockObject("plain_text", "Next", false, false)
	nextBtn := slack.NewButtonBlockElement("", "next_3", nextBtnTxt)

	prevBtnTxt := slack.NewTextBlockObject("plain_text", "Prve", false, false)
	prevBtn := slack.NewButtonBlockElement("", "prev_2", prevBtnTxt)

	actionBlock := slack.NewActionBlock("", prevBtn, nextBtn)

	bookListBlocks = append(bookListBlocks, actionBlock)

	message := slack.NewBlockMessage(
		bookListBlocks...,
	)

	message.ReplaceOriginal = true

	return json.MarshalIndent(message, "", "    ")
}
