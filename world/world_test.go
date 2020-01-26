package world

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var world = World{}

func init() {
	world = create(10, 10)
}

func TestCreate(t *testing.T) {
	var message = Create{5, 5}

	world = New(message)

	assert.Equal(t, 5, len(world.Cells))
}


func TestModify(t *testing.T) {
	var message = Modify{Position{1, 1}, Wall}

	world = world.Update(message)

	assert.Equal(t, Wall, world.Cells[1][1])
}
