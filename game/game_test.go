package main

import (
	c "alive-dungeon/creature"
	i "alive-dungeon/interactions"
	w "alive-dungeon/world"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGame(t *testing.T) {
	var world = w.New(w.Create{6, 6})
	world.Update(w.Modify{Position: w.Position{5, 5}, NewType: w.Wall})
	world.Update(w.Modify{Position: w.Position{2, 5}, NewType: w.Wall})
	world.Update(w.Modify{Position: w.Position{5, 3}, NewType: w.Wall})
	world.Update(w.Modify{Position: w.Position{4, 2}, NewType: w.Wall})

	var initialState = State{world,
		[]c.Creature{
			c.New(c.Create{Position: c.Position{1, 1}}),
			c.New(c.Create{Position: c.Position{2, 2}}),
			c.New(c.Create{Position: c.Position{3, 3}}),
		}}

	state := make(chan State, 1)
	state <- initialState
	actions := make(chan c.Direction)
	messages := make(chan string, 1)
	messages <- Print(initialState.World, initialState.Creatures)

	go input(actions)
	go process(0, state, actions, messages)
	go process(1, state, actions, messages)
	go process(2, state, actions, messages)

	output(messages)
}

func input(actions chan c.Direction) {
	rand.Seed(time.Now().Unix())
	for {
		var action = rand.Intn(4)
		switch action {
		case 0:
			actions <- c.Up
		case 1:
			actions <- c.Down
		case 2:
			actions <- c.Left
		case 3:
			actions <- c.Right
		}
		time.Sleep(time.Second)
	}
}

func output(messages <-chan string) {
	for {
		printMessage(<-messages)
	}
}

func printMessage(message string) {
	for i := 0; i < 10; i++ {
		fmt.Print("~")
	}
	fmt.Println()
	fmt.Print(message)
	fmt.Println()
	for i := 0; i < 10; i++ {
		fmt.Print("~")
	}
	fmt.Println()
}

func process(creatureId int, states chan State, movement <-chan c.Direction, messages chan<- string) {
	for {
		var message = ""
		var state = <-states
		var direction = <-movement
		move := c.Move{Direction: direction}
		var result = i.Movement(state.Creatures[creatureId], state.World, state.Creatures, move)
		if result == i.Success {
			state.Creatures[creatureId] = state.Creatures[creatureId].Update(move)
			message += "Te moviste hacia " + parseDirection(direction)
		} else {
			message += "No pudiste moverte hacia " + parseDirection(direction)
		}
		message += "\n\n" + Print(state.World, state.Creatures)
		messages <- message
		states <- state
	}
}

func parseDirection(direction c.Direction) string {
	switch direction {
	case c.Up:
		return "arriba"
	case c.Down:
		return "abajo"
	case c.Left:
		return "izquierda"
	case c.Right:
		return "derecha"
	}
	return ""
}

//only for visual testing
func Print(world w.World, creatures []c.Creature) string {
	var width = len(world.Cells)
	var height = len(world.Cells[0])
	var message = make([][]string, width)
	for x := 0; x < width; x++ {
		message[x] = make([]string, height)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			cell := world.Cells[y][x]
			switch cell {
			case w.Wall:
				message[y][x] = "#"
			case w.Empty:
				message[y][x] = "·"
			}
		}
	}
	for _, creature := range creatures {
		message[creature.Position.Y][creature.Position.X] = "r"
	}

	var result = ""
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			result += message[y][x]
		}
		result += "\n"
	}
	return result
}
