package game

import (
	"testing"
	w "alive-dungeon/world"
	c "alive-dungeon/creature"
)

func TestAddCreature(t *testing.T) {
	walls := []w.Position{
		{1, 1},
		{2, 2},
		{2, 1},
	}
	request := w.CreationRequest{3, 3, walls}
	world := w.Create(request)
	creature := c.Creature{}

	AddCreature(world, creature, w.Position{2, 3})

	//assert
}

func AddCreature(world w.World, creature c.Creature, position w.Position) {

}
