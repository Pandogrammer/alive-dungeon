package world

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreation(t *testing.T) {
	request := CreationRequest{Width: 2, Height: 5}

	result := Create(request)

	assert.Equal(t, 2, len(result.Cells))
	assert.Equal(t, 5, len(result.Cells[0]))
}

func TestCreationWithWalls(t *testing.T) {
	request := CreationRequest{Width: 3, Height: 3, Walls: []Position{{1, 1}}}
	world := Create(request)

	assert.Equal(t, Wall, world.Cells[1][1])
}

func TestAddWall(t *testing.T) {
	request := CreationRequest{Width: 3, Height: 3}
	world := Create(request)
	world.AddWall(Position{1, 1})

	assert.Equal(t, Wall, world.Cells[1][1])
}
