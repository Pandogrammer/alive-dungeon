package world

type Modify struct {
	Position Position
	NewType  Cell
}

type Create struct {
	Width  int
	Height int
}
