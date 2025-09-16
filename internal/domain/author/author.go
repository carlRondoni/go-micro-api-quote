package author

import (
	"go-micro-api-quote/internal/domain/author/author_value_objects"

	"github.com/google/uuid"
)

type Author struct {
	Id   author_value_objects.Id
	Name author_value_objects.Name
}

func NewAuthorFromString(name string) Author {
	return Author{
		Id:   author_value_objects.NewIdFromUuid(uuid.New()),
		Name: author_value_objects.NewNameFromString(name),
	}
}
