package iso3166_test

import (
	"testing"

	"github.com/foomo/gostandards/iso3166"
	"github.com/stretchr/testify/assert"
)

func TestAlpha2(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "DE", iso3166.Alpha2DE.String())
}

func TestAlpha2Valid(t *testing.T) {
	t.Parallel()

	assert.True(t, iso3166.Alpha2DE.Valid())
	assert.True(t, iso3166.Alpha2("DE").Valid())
	assert.False(t, iso3166.Alpha2("").Valid())
	assert.False(t, iso3166.Alpha2("XX").Valid())
}
