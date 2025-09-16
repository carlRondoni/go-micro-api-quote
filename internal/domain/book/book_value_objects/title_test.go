package book_value_objects_test

import (
	"go-micro-api-quote/internal/domain/book/book_value_objects"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
)

func TestTitle(t *testing.T) {
	title := gofakeit.Word()

	ttl := book_value_objects.NewTitleFromString(title)

	assert.NotNil(t, ttl)
	assert.Equal(t, title, ttl.Value())
}
