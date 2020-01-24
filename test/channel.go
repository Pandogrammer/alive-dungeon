package main

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}

type Action int

const (
	Increment Action = iota
	Decrement
)

func main() {
	actions := make(chan Action, 1)
	changes := make(chan int, 1)
	initialState := 0

	process(initialState, actions, changes)
	sender(actions)
	receiver(changes)
}

func receiver(changes chan int) {
	for {
		fmt.Println(<-changes)
	}
}

func sender(actions chan Action) {
	go func() {
		for {
			actions <- Increment
			time.Sleep(time.Second)
		}
	}()
}

func process(initialState int, actions <-chan Action, changes chan<- int) {
	go func() {
		var state = initialState
		for {
			switch <-actions {
			case Increment:
				state = state + 1
				changes <- state
			}
		}
	}()
}
