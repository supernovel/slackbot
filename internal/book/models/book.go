package book

import (
	"time"
)

// Book : bind book document
type Book struct {
	ISBN            string
	Title           string
	Author          string
	Publisher       string
	PublicationDate time.Time
}
