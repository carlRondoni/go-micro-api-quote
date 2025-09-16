package book

import (
	"go-micro-api-quote/internal/domain/author"
	"go-micro-api-quote/internal/domain/author/author_value_objects"
	"go-micro-api-quote/internal/domain/book/book_value_objects"

	"github.com/google/uuid"
)

type Book struct {
	Id     book_value_objects.Id
	Title  book_value_objects.Title
	Author author_value_objects.Name
}

func NewBook(title string, athr author.Author) Book {
	return Book{
		Id:     book_value_objects.NewIdFromUuid(uuid.New()),
		Title:  book_value_objects.NewTitleFromString(title),
		Author: athr.Name,
	}
}
