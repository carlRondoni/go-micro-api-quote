package author_test

import (
	"go-micro-api-quote/internal/domain/author"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	name := gofakeit.FirstName()

	nameObj := author.NewNameFromString(name)

	assert.Equal(t, name, nameObj.Value())
}
