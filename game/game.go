package game

import (
	c "alive-dungeon/creature"
	w "alive-dungeon/world"
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
