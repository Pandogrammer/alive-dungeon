package world

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreation(t *testing.T) {
	request := CreationRequest{width: 2, height: 5}

	result := Create(request)

	assert.Equal(t, 2, len(result.cells))
	assert.Equal(t, 5, len(result.cells[0]))
}

func TestCreationWithWalls(t *testing.T) {
	request := CreationRequest{3, 3, []Position{{1, 1}}}
	world := Create(request)

	_, result := world.cells[1][1].(Wall)

	assert.True(t, result)
}

func TestPrintMap(t *testing.T) {
	walls := []Position{
		{1, 1},
		{2, 2},
		{2, 1},
	}
	request := CreationRequest{3, 3, walls}
	world := Create(request)

	PrintMap(world)
}

//only for visual testing
func PrintMap(world World) {
	var lastIndex = 0
	for line, cell := range world.cells {
		if line != lastIndex {
			fmt.Print("\n")
			lastIndex = line
		}

		for _, x := range cell {
			switch x.(type) {
			case Wall:
				fmt.Printf("#")
			default:
				fmt.Printf("·")
			}
		}
	}
	fmt.Println()
}
