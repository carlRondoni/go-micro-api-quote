package quote_value_objects_test

import (
	"go-micro-api-quote/internal/domain/quote/quote_value_objects"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestId(t *testing.T) {
	t.Run("create an id using uuid", func(t *testing.T) {
		value := uuid.New()
		vo := quote_value_objects.NewIdFromUuid(value)

		assert.NotNil(t, vo)
		assert.Equal(t, value, vo.Value())
	})

	t.Run("create an id using string uuid", func(t *testing.T) {
		value := gofakeit.UUID()
		vo, err := quote_value_objects.NewIdFromString(value)
		assert.NotNil(t, vo)
		assert.NoError(t, err)

		assert.Equal(t, value, vo.String())
	})

	t.Run("error on wrong string uuid", func(t *testing.T) {
		value := gofakeit.Word()
		_, err := quote_value_objects.NewIdFromString(value)
		assert.Error(t, err)
	})
}
