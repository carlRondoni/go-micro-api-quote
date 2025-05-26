package quote

import (
	"go-micro-api-quote/internal/author"
	"go-micro-api-quote/internal/book"

	"github.com/google/uuid"
)

type Quote struct {
	Id     Id
	Text   Text
	Author author.Author
	Book   book.Book
}

func NewFromValues(
	athr string,
	txt string,
) Quote {
	return Quote{
		Id:     NewIdFromUuid(uuid.New()),
		Text:   NewtextFromString(txt),
		Author: author.NewAuthorFromString(athr),
	}
}
