package book

import (
	"testing"

	book "github.com/supernovel/slackbot/internal/book/models"
)

// TODO. wantMesage
// Check block body data.
//
// Check block action button.
//
// current == 0     => show next button
//
// current == total => show prev button
//
// current < total  => show next, prev button
//
// current > total  => not possible
func TestBooks_BuildBookListBlock(t *testing.T) {
	tests := []struct {
		book        []book.Book
		current     int
		total       int
		wantMessage string
	}{
		{
			book: []book.Book{
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
			},
			current:     0,
			total:       2,
			wantMessage: "",
		},
		{
			book: []book.Book{
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
			},
			current:     2,
			total:       2,
			wantMessage: "",
		},
		{
			book: []book.Book{
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
			},
			current:     2,
			total:       4,
			wantMessage: "",
		},
	}

	for i, test := range tests {
		message, err := BuildBookListBlock(&test.book, test.current, test.total)

		if err != nil {
			t.Fatalf("%d: Unexpected error: %s", i, err)
		}

		if string(message) != test.wantMessage {
			t.Errorf("%d: Got message %s, want %s", i, message, test.wantMessage)
		}
	}
}
