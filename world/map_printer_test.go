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
		{6, 1},
		{4, 8},
	}
	request := CreationEvent{Width: 10, Height: 10, Walls: walls}
	world := Create(request)

	PrintMap(world)
}

//only for visual testing
func PrintMap(world World) {
	var lastIndex = 0
	for line, column := range world.Cells {
		if line != lastIndex {
			fmt.Print("\n")
			lastIndex = line
		}

		for _, cell := range column {
			switch cell {
			case Wall:
				fmt.Printf("#")
			case Empty:
				fmt.Printf(".")
			}
		}
	}
	fmt.Println()
}
