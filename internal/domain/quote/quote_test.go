package quote_test

import (
	"go-micro-api-quote/internal/domain/quote"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestQuote(t *testing.T) {
	auth := gofakeit.FirstName()
	txt := gofakeit.Quote()

	qt := quote.NewFromValues(
		auth,
		txt,
	)

	assert.Equal(t, auth, qt.Author.Name.Value())
	assert.Equal(t, txt, qt.Text.Value())
	assert.NotEqual(t, uuid.Nil, qt.Id.Value())
}
