import (
	"time"
)

type Book struct {
	ISBN string
	Title string
	Author string
	Publisher string
	PublicationDate time.Time
}