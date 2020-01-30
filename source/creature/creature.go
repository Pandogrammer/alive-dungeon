package creature

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Creature struct {
	Position Position
	Health   int
}

func (c Creature) Damage(value int) Creature {
	c.Health = c.Health - value
	return c
}
type Position struct {
	X int
	Y int
}
