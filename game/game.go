package main

import (
	c "alive-dungeon/creature"
	w "alive-dungeon/world"
)

type State struct {
	World    w.World
	Creatures []c.Creature
}
