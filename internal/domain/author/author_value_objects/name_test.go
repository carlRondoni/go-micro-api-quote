package author_value_objects_test

import (
	"go-micro-api-quote/internal/domain/author/author_value_objects"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	name := gofakeit.FirstName()

	nameObj := author_value_objects.NewNameFromString(name)

	assert.NotNil(t, nameObj)
	assert.Equal(t, name, nameObj.Value())
}
