package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Action int

const (
	Increment Action = iota
	Decrement
)

func main() {
	actions := make(chan Action)
	state := make(chan int)
	initialState := 0

	go process(initialState, actions, state)
	go sender(actions, Increment)
	go sender(actions, Decrement)
	go sender(actions, Increment)
	go sender(actions, Decrement)
	go input(actions)

	output(state)
}

func input(actions chan Action) {
	reader := bufio.NewReader(os.Stdin)
	for {
		message, _ := reader.ReadString('\n')
		if message == "s\n" {
			actions <- Increment
		}
		if message == "d\n" {
			actions <- Decrement
		}
	}
}

func output(state <-chan int) {
	var counter = 0
	for {
		newState := <-state
		fmt.Print("Actual state: ")
		fmt.Println(newState)
		switch newState {
		default:
			counter++
			if counter > 3 {
				fmt.Print("Last state: ")
				fmt.Println(newState)
				counter = 0
			}
		}
	}
}

func sender(actions chan Action, action Action) {
	for {
		actions <- action
		time.Sleep(time.Second)
	}
}

func process(initialState int, actions <-chan Action, state chan<- int) {
	var actualState = initialState
	for {
		switch <-actions {
		case Increment:
			actualState = actualState + 1
			state <- actualState
		case Decrement:
			actualState = actualState - 1
			state <- actualState
		}
	}
}
