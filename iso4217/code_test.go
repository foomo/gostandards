package iso4217_test

import (
	"testing"

	"github.com/foomo/gostandards/iso4217"
	"github.com/stretchr/testify/assert"
)

func TestValid(t *testing.T) {
	t.Parallel()

	assert.True(t, iso4217.CodeEUR.Valid())
	assert.True(t, iso4217.Code("EUR").Valid())
	assert.False(t, iso4217.Code("").Valid())
	assert.False(t, iso4217.Code("FOO").Valid())
}

func TestString(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "EUR", iso4217.CodeEUR.String())
	assert.Equal(t, "EUR", iso4217.Code("EUR").String())
}
