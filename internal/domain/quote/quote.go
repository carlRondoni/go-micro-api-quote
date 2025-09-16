package quote

import (
	"go-micro-api-quote/internal/domain/author"
	"go-micro-api-quote/internal/domain/quote/quote_value_objects"

	"github.com/google/uuid"
)

type Quote struct {
	Id     quote_value_objects.Id
	Text   quote_value_objects.Text
	Author author.Author
}

func NewFromValues(
	athr string,
	txt string,
) Quote {
	return Quote{
		Id:     quote_value_objects.NewIdFromUuid(uuid.New()),
		Text:   quote_value_objects.NewTextFromString(txt),
		Author: author.NewAuthorFromString(athr),
	}
}
