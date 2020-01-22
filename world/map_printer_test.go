package world

import (
	"fmt"
	"testing"
)

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
	for line, cell := range world.Cells {
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