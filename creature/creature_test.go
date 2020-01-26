package creature

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	var message = Create{Position{5, 5}}

	var creature = New(message)

	assert.Equal(t, Position{5, 5}, creature.Position)
}

func TestMove(t *testing.T) {
	var creature = New(Create{Position{5, 5}})

	var rightMessage = Move{Right}
	var right = creature.Update(rightMessage)

	var leftMessage = Move{Left}
	var left = creature.Update(leftMessage)

	var upMessage = Move{Up}
	var up = creature.Update(upMessage)

	var downMessage = Move{Down}
	var down = creature.Update(downMessage)

	assert.Equal(t, Position{6, 5}, right.Position)
	assert.Equal(t, Position{4, 5}, left.Position)
	assert.Equal(t, Position{5, 4}, up.Position)
	assert.Equal(t, Position{5, 6}, down.Position)
}
