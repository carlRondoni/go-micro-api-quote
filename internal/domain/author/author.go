package author

import "github.com/google/uuid"

type Author struct {
	Id   Id
	Name Name
}

func NewAuthorFromString(name string) Author {
	return Author{
		Id:   NewIdFromUuid(uuid.New()),
		Name: NewNameFromString(name),
	}
}
