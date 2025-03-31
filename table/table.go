package table

import (
	"fmt"

	"github.com/andrewbruno/go-toyrobot/compass"
	"github.com/andrewbruno/go-toyrobot/toy"
	"github.com/andrewbruno/go-toyrobot/unit"
)

var (
	_tableLengthX unit.X // Value set at initialization
	_tableLengthY unit.Y // Value set at initialization
)

const (
	msgRobotNotPlaced = "Robot has not been placed"
)

// Tabler is the interface defining functions that must be implemented
type Tabler interface {
	Place(unit.X, unit.Y, compass.Direction) *toy.Robot
	Move()
	Left()
	Right()
	Report() string
}

// Table represents the object (table top) where the Toy will be placed on
type Table struct {
	lengthX unit.X
	lengthY unit.Y
	robot   *toy.Robot
}

// NewTable returns a table object.
func NewTable() *Table {
	return &Table{lengthX: _tableLengthX, lengthY: _tableLengthY}
}

// Place puts the robot in the located x,y stated, facing the direction f
func (t *Table) Place(x unit.X, y unit.Y, f compass.Direction) *toy.Robot {

	if !t.moveAllowed(x, y) {
		return t.robot // Could be nil or set
	}

	if t.robot == nil {
		t.robot = toy.NewRobot()
	}

	t.robot.Update(x, y, f)
	return t.robot
}

// Report returns the location of the robot.
func (t *Table) Report() string {
	if t.robot != nil {
		// Work out the next x,y,f
		return fmt.Sprintf("%d,%d,%s", t.robot.X, t.robot.Y, t.robot.F)
	}
	return msgRobotNotPlaced
}

// Move the robot IFF it has been placed.
func (t *Table) Move() {
	if t.robot != nil {

		x, y := t.robot.X, t.robot.Y

		switch t.robot.F {
		case compass.NORTH:
			y++
		case compass.EAST:
			x++
		case compass.SOUTH:
			y--
		case compass.WEST:
			x--
		}

		if t.moveAllowed(x, y) {
			t.robot.X, t.robot.Y = x, y
		}

	}
}

// Left turns the robot counter-clockwise 90 degrees IFF it has been placed.
func (t *Table) Left() {
	if t.robot != nil {
		t.robot.F = compass.Left(t.robot.F)
	}
}

// Right turns the robot 90 degres clockwise IFF it has been placed.
func (t *Table) Right() {
	if t.robot != nil {
		t.robot.F = compass.Right(t.robot.F)
	}
}

// moveAllowed checks to see if the robot is in bounds or not.
func (t *Table) moveAllowed(x unit.X, y unit.Y) bool {
	if x >= t.lengthX || y >= t.lengthY {
		return false
	}
	return true
}

func init() {
	// Can be extended via a config file
	// Must be > 0 and less than 255
	_tableLengthX = unit.X(5)
	_tableLengthY = unit.Y(5)
}
