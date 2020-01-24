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
	actions := make(chan Action, 1)
	state := make(chan int, 1)
	initialState := 0

	go process(initialState, actions, state)
	go sender(actions)
	go input(actions)

	receiver(state)
}

func input(actions chan Action) {
	reader := bufio.NewReader(os.Stdin)
	for {
		message, _ := reader.ReadString('\n')
		if message == "x\n" {
			actions <- Decrement
		}
	}
}

func receiver(state chan int) {
	for {
		fmt.Println(<-state)
	}
}

func sender(actions chan Action) {
	for {
		actions <- Increment
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
