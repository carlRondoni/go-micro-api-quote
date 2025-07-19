package book_test

import (
	"go-micro-api-quote/internal/domain/book"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
)

func TestTitle(t *testing.T) {
	title := gofakeit.Word()

	ttl := book.NewTitleFromString(title)

	assert.Equal(t, title, ttl.Value())
}
