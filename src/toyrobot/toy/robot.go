package toy

import (
	"toyrobot/compass"
	"toyrobot/unit"
)

// Robot object
type Robot struct {
	x unit.X
	y unit.Y
	f compass.Direction
}

// NewRobot returns a new Robot object
func NewRobot() *Robot {
	return &Robot{}
}

// Update updates the location of the robot.
func (r *Robot) Update(x unit.X, y unit.Y, f compass.Direction) {
	r.x = x
	r.y = y
	r.f = f
}

// X returns y location
func (r *Robot) X() unit.X {
	return r.x
}

// Y returns y location
func (r *Robot) Y() unit.Y {
	return r.y
}

// F returns the compass direction robot is facing
func (r *Robot) F() compass.Direction {
	return r.f
}
