package author_test

import (
	"go-micro-api-quote/internal/author"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
)

func TestAuthor(t *testing.T) {
	athrName := gofakeit.Name()
	authorObj := author.NewAuthorFromString(athrName)

	assert.Equal(t, athrName, authorObj.Name.Value())
	assert.NotNil(t, authorObj.Id)
}
