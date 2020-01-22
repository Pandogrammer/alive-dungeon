package creature

import (
	w "alive-dungeon/world"
	"testing"
)

var world

func init(){
	request := w.CreationRequest{Width: 3, Height: 3, Walls: []w.Position{{1, 1}}}
	world = w.Create(request)
}

func TestWorldInteraction(t *testing.T) {
	rat := Create()
}


type Rat Creature

func Create() Rat {
	return Rat{}
}