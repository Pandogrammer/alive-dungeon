package game

import (
	c "alive-dungeon/creature"
	w "alive-dungeon/world"
)

type Game struct {
	World     w.World
	Creatures []SpawnedCreature
}

type SpawnedCreature struct {
	Creature c.Creature
	Position w.Position
}

func (game Game) AddCreature(creature SpawnedCreature) Game {
	game.Creatures = append(game.Creatures, creature)
	return game
}

func Create(world w.World) Game {
	return Game{World: world, Creatures: []SpawnedCreature{}}
}
