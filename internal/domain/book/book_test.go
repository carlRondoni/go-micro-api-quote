package book_test

import (
	"go-micro-api-quote/internal/author"
	"go-micro-api-quote/internal/book"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
)

func TestBook(t *testing.T) {
	title := gofakeit.Name()

	athrObj := author.NewAuthorFromString(gofakeit.FirstName())

	bk := book.NewBook(title, athrObj)

	assert.Equal(t, title, bk.Title.Value())
	assert.Equal(t, athrObj.Name, bk.Author)
}
