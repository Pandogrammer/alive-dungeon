package interactions

import (
	c "alive-dungeon/creature"
	w "alive-dungeon/world"
	"math"
)

func HasCreature(creatures []c.Creature, position c.Position) bool {
	i := 0
	hasCreature := false
	for i < len(creatures) && !hasCreature {
		if creatures[i].Position == position {
			hasCreature = true
		}
		i++
	}
	return hasCreature
}

type Sight map[c.Position]Representation

type Representation int

func (r Representation) ToString() string {
	switch r {
	case Fog:
		return "~"
	case Wall:
		return "#"
	case Empty:
		return "·"
	case Creature:
		return "r"
	case Void:
		return "%"
	}
	return "X"
}

const (
	Fog Representation = iota
	Wall
	Empty
	Creature
	Void
)

type View [][]Representation

func (v View) ToString() string {
	var size = len(v)
	var result = ""
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			result += v[i][j].ToString()
		}
		result += "\n"
	}
	return result
}

type State struct {
	World     w.World
	Creatures []c.Creature
}

func FullVision(creature c.Creature, state State, distance int) View {
	var sight = distance*2 + 1
	var vision = make([][]Representation, sight)
	for i := 0; i < sight; i++ {
		vision[i] = make([]Representation, sight)
		for j := 0; j < sight; j++ {
			vision[i][j] = Fog
			var pos = c.Position{
				X: j + creature.Position.X - distance,
				Y: i + creature.Position.Y- distance}
			var onSight = IsOnSight(creature.Position, pos, distance)
			if onSight {
				vision[i][j] = getRepresentation(state, pos)
			}
		}
	}
	return vision
}

func IsOnSight(position c.Position, pos c.Position, distance int) bool {
	return math.Sqrt(math.Pow(float64(position.X-pos.X), 2)+math.Pow(float64(position.Y-pos.Y), 2)) <= float64(distance)
}

func Vision(creature c.Creature, state State, distance int) Sight {
	var position = creature.Position
	var sight = make(Sight)
	var representation Representation
	for _, pos := range positionsInSight(position, distance) {
		representation = getRepresentation(state, pos)
		sight[pos] = representation
	}
	return sight
}

func getRepresentation(state State, pos c.Position) Representation {
	if state.World.OutOfBounds(w.Position(pos)) {
		return Void
	} else if HasCreature(state.Creatures, pos) {
		return Creature
	} else {
		switch state.World.CellAt(pos.X, pos.Y) {
		case w.Wall:
			return Wall
		case w.Empty:
			return Empty
		}
	}
	return Void
}

func positionsInSight(position c.Position, distance int) []c.Position {
	var vision []c.Position

	var j = 0
	var minX = position.X - distance
	var maxX = position.X + distance
	for i := minX; i <= maxX; i++ {
		for y := -j; y < 0; y++ {
			var posY = position.Y + y
			vision = append(vision, c.Position{X: i, Y: posY})
		}

		vision = append(vision, c.Position{X: i, Y: position.Y})

		for y := j; y > 0; y-- {
			var posY = position.Y + y
			vision = append(vision, c.Position{X: i, Y: posY})
		}

		if i < position.X {
			j++
		} else {
			j--
		}
	}

	return vision
}
