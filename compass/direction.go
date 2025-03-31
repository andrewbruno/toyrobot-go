package compass

import "fmt"

// Direction is a custom type representing the direction the toy is facing.
type Direction uint8

// The list (enums) of possible directions the toy may face
// Using iota, so 0, 1, 2, 3 is actual value going clockwise starting from NORTH
// in 90 degree movements only.
const (
	NORTH Direction = iota // 0
	EAST                   // 1
	SOUTH                  // 2
	WEST                   // 3
)

var mapDirections = make(map[string]Direction)

// Left returns the compass direction when currentDirection is rotated
// 90 degrees counter-clockwise.
func Left(currentDirection Direction) Direction {
	if currentDirection == NORTH {
		// Edge case
		return WEST
	}

	return currentDirection - 1
}

// Right returns the compass direction when currentDirection is rotated
// 90 degrees clockwise.
func Right(currentDirection Direction) Direction {
	if currentDirection == WEST {
		// Edge case
		return NORTH
	}

	return currentDirection + 1
}

// Returns the string representation of the constant name.
func (d Direction) String() string {
	var direction string

	switch d {
	case NORTH:
		direction = "NORTH"
	case EAST:
		direction = "EAST"
	case SOUTH:
		direction = "SOUTH"
	case WEST:
		direction = "WEST"
	}

	return direction
}

// ParseDirection interprets a string s and returns the corresponding value.
func ParseDirection(s string) (*Direction, error) {
	direction, exists := mapDirections[s]
	if !exists {
		return nil, fmt.Errorf("Can not parse direction [%s]", s)
	}

	return &direction, nil
}

func init() {
	mapDirections["NORTH"] = NORTH
	mapDirections["EAST"] = EAST
	mapDirections["SOUTH"] = SOUTH
	mapDirections["WEST"] = WEST
}
