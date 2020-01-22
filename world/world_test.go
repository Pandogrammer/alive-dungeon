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
	request := CreationRequest{3, 3, []Position{{1, 1}}}
	world := Create(request)

	_, result := world.Cells[1][1].(Wall)

	assert.True(t, result)
}

func TestAddWall(t *testing.T) {
	request := CreationRequest{3, 3, nil}
	world := Create(request)

	_, result := world.Cells[1][1].(Wall)

	assert.True(t, result)
}
