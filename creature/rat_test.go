package creature

import (
	w "alive-dungeon/world"
	"testing"
)

var world

func init(){
	request := w.CreationRequest{3, 3, []w.Position{{1, 1}}}
	world = w.Create(request)
}

func TestWorldInteraction(t *testing.T) {
	rat := Create()
}


type Rat Creature

func Create() Rat {
	return Rat{}
}