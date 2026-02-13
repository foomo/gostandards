package iso3166_test

import (
	"testing"

	"github.com/foomo/gostandards/iso3166"
	"github.com/stretchr/testify/assert"
)

func TestAlpha3(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "DEU", iso3166.Alpha3DEU.String())
}

func TestAlpha3Valid(t *testing.T) {
	t.Parallel()

	assert.True(t, iso3166.Alpha3DEU.Valid())
	assert.True(t, iso3166.Alpha3("DEU").Valid())
	assert.False(t, iso3166.Alpha3("").Valid())
	assert.False(t, iso3166.Alpha3("XXX").Valid())
}
