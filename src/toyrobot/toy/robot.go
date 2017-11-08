package toy

import (
	"toyrobot/compass"
	"toyrobot/unit"
)

// Robot object
type Robot struct {
	X unit.X
	Y unit.Y
	F compass.Direction
}

// NewRobot returns a new Robot object
func NewRobot() *Robot {
	return &Robot{}
}

// Update updates the location of the robot.
func (r *Robot) Update(x unit.X, y unit.Y, f compass.Direction) {
	r.X = x
	r.Y = y
	r.F = f
}
