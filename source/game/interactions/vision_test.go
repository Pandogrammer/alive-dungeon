package interactions

import (
	c "alive-dungeon/source/creature"
	w "alive-dungeon/source/world"
	"fmt"
	"testing"
)

type World struct {
	Width int
	Height int
}

func TestFullVision(t *testing.T) {
	var world = w.New(w.Create{
		Width:  10,
		Height: 10,
	})
	var creature = c.New(struct{ Position c.Position }{Position: struct {
		X int
		Y int
	}{X: 7, Y: 7}})
	var state = State{
		World:     world,
		Creatures: []c.Creature{creature},
	}

	var vision = FullVision(creature, state, 4).ToString()

	fmt.Print(vision)
}

