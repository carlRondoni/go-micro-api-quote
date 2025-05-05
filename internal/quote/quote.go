package quote

import "github.com/google/uuid"

type Quote struct {
	Id     Id
	Text   Text
	Author Author
}

func NewFromValues(
	author string,
	txt string,
) Quote {
	return Quote{
		Id:     NewIdFromUuid(uuid.New()),
		Text:   NewtextFromString(txt),
		Author: NewAuthorFromString(author),
	}
}
