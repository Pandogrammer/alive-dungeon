package clock

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTime(t *testing.T) {
	timer := Clock{}

	timer = timer.Advance()
	timer = timer.Advance()

	assert.Equal(t, 2, timer.Actual)
}
