package book

import (
	"go-micro-api-quote/internal/author"

	"github.com/google/uuid"
)

type Book struct {
	Id     Id
	Title  Title
	Author author.Name
}

func NewBook(title string, athr author.Author) Book {
	return Book{
		Id:     NewIdFromUuid(uuid.New()),
		Title:  NewTitleFromString(title),
		Author: athr.Name,
	}
}
