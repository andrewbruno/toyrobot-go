package toy

import (
	"toyrobot/compass"
	"toyrobot/unit"
)

// Toyer definition
type Toyer interface {
	Update(unit.X, unit.Y, compass.Direction)
	X() unit.X
	Y() unit.Y
	F() compass.Direction
}
