package quote_value_objects_test

import (
	"go-micro-api-quote/internal/domain/quote/quote_value_objects"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
)

func TestText(t *testing.T) {
	value := gofakeit.Quote()
	vo := quote_value_objects.NewTextFromString(value)

	assert.NotNil(t, vo)
	assert.Equal(t, value, vo.Value())
}
