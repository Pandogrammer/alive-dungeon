package game

import (
	c "alive-dungeon/source/creature"
	w "alive-dungeon/source/world"
)

type State struct {
	World     w.World
	Creatures []c.Creature
}

type Connection struct {
	Id       int
	State    chan State
	Actions  chan c.Direction
	Messages chan string
}
