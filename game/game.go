package main

import (
	c "alive-dungeon/creature"
	i "alive-dungeon/interactions"
	w "alive-dungeon/world"
	"bufio"
	"fmt"
	"os"
)

type State struct {
	World    w.World
	Creature c.Creature
}

func main() {
	var world = w.New(w.Create{6, 6})
	world.Update(w.Modify{Position: w.Position{5, 5}, NewType: w.Wall})
	world.Update(w.Modify{Position: w.Position{2, 5}, NewType: w.Wall})
	world.Update(w.Modify{Position: w.Position{5, 3}, NewType: w.Wall})
	world.Update(w.Modify{Position: w.Position{4, 2}, NewType: w.Wall})
	var creature = c.New(c.Create{Position: c.Position{2, 2}})
	var initialState = State{world, creature}

	state := make(chan State, 1)
	state <- initialState
	actions := make(chan c.Direction)
	messages := make(chan string, 1)
	messages <- Print(initialState.World, initialState.Creature)

	go input(actions)
	go process(state, actions, messages)

	output(messages)
}

func input(actions chan c.Direction) {
	reader := bufio.NewReader(os.Stdin)
	for {
		message, _ := reader.ReadString('\n')
		if message == "w\n" {
			actions <- c.Up
		}
		if message == "a\n" {
			actions <- c.Left
		}
		if message == "s\n" {
			actions <- c.Down
		}
		if message == "d\n" {
			actions <- c.Right
		}
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

func process(states chan State, movement <-chan c.Direction, messages chan<- string) {
	for {
		var message = ""
		var state = <-states
		move := c.Move{Direction: <-movement}
		var result = i.Movement(state.World, state.Creature, move)
		if result == i.Success {
			state.Creature = state.Creature.Update(move)
			message += "Te moviste"
		} else {
			message += "No pudiste moverte"
		}
		message += "\n\n" + Print(state.World, state.Creature)
		messages <- message
		states <- state
	}
}

//only for visual testing
func Print(world w.World, creature c.Creature) string {
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

	message[creature.Position.Y][creature.Position.X] = "r"

	var result = ""
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			result += message[y][x]
		}
		result += "\n"
	}
	return result
}
