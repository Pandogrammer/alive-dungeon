package game

import (
	c "alive-dungeon/creature"
	w "alive-dungeon/world"
)

type Game struct {
	World     w.World
	Creatures []c.Creature
}

func (game Game) AddCreature(creature c.Creature, position w.Position) Game {
	game.Creatures = append(game.Creatures, creature)
	return game
}

func Create(world w.World) Game {
	return Game{World: world, Creatures: []c.Creature{}}
}
