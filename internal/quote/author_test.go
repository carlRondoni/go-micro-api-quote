package quote_test

import (
	"go-micro-api-quote/internal/quote"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
)

func TestAuthor(t *testing.T) {
	author := gofakeit.Name()
	authorVo := quote.NewAuthorFromString(author)

	assert.Equal(t, author, authorVo.Value())
}
