package quote_test

import (
	"go-micro-api-quote/internal/domain/quote"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
)

func TestText(t *testing.T) {
	value := gofakeit.Quote()
	vo := quote.NewtextFromString(value)

	assert.Equal(t, value, vo.Value())
}
