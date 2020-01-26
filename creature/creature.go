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
}
type Position struct {
	X int
	Y int
}
