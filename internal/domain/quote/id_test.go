package quote_test

import (
	"go-micro-api-quote/internal/quote"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestId(t *testing.T) {
	t.Run("create an id using uuid", func(t *testing.T) {
		value := uuid.New()
		vo := quote.NewIdFromUuid(value)

		assert.Equal(t, value, vo.Value())
	})

	t.Run("create an id using string uuid", func(t *testing.T) {
		value := gofakeit.UUID()
		vo, err := quote.NewIdFromString(value)
		assert.NoError(t, err)

		assert.Equal(t, value, vo.String())
	})

	t.Run("error on wrong string uuid", func(t *testing.T) {
		value := gofakeit.Word()
		_, err := quote.NewIdFromString(value)
		assert.Error(t, err)
	})
}
