package query

import (
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogin(t *testing.T) {
	var s *LoginQuery
	err := validator.New().Struct(s)
	assert.Nil(t, err, err)
}
