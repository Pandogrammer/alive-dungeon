package interactions

import (
	c "alive-dungeon/creature"
	g "alive-dungeon/game"
	w "alive-dungeon/world"
	"github.com/stretchr/testify/assert"
	"testing"
)

var creature = c.Creature{}
var world = w.World{}
var state = g.State{}

type Position struct {
	X int
	Y int
}

func TestVision(t *testing.T) {
	givenAWorld()

	var result = Vision(creature, state)
	var expected = []Position{
		{X: 0, Y: 2},
		{X: 1, Y: 1},
		{X: 1, Y: 2},
		{X: 1, Y: 3},
		{X: 2, Y: 0},
		{X: 2, Y: 1},
		{X: 2, Y: 2},
		{X: 2, Y: 4},
		{X: 2, Y: 3},
		{X: 3, Y: 1},
		{X: 3, Y: 2},
		{X: 3, Y: 3},
		{X: 4, Y: 2},
	}

	assert.Equal(t, expected, result)
}

func Vision(creature c.Creature, s g.State) []Position {
	var distance = 2
	var position = creature.Position
	return positionsInSight(position, distance)
}

func positionsInSight(position c.Position, distance int) []Position {
	var vision []Position

	var j = 0
	var minX = position.X - distance
	var maxX = position.X + distance
	for i := minX; i <= maxX; i++ {
		for y := -j; y < 0; y++ {
			var posY = position.Y + y
			vision = append(vision, Position{X: i, Y: posY})
		}

		vision = append(vision, Position{X: i, Y: position.Y})

		for y := j; y > 0; y-- {
			var posY = position.Y + y
			vision = append(vision, Position{X: i, Y: posY})
		}

		if i < position.X {
			j++
		} else {
			j--
		}
	}

	return vision
}

func givenAWorld() {
	creature = c.New(struct{ Position c.Position }{Position: struct {
		X int
		Y int
	}{X: 2, Y: 2}})

	world = w.New(struct {
		Width  int
		Height int
	}{Width: 10, Height: 10})

	state = g.State{
		World:     world,
		Creatures: []c.Creature{creature},
	}
}
