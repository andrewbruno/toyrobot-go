package table

import (
	"fmt"
	"testing"

	"github.com/andrewbruno/go-toyrobot/src/toyrobot/compass"
	"github.com/andrewbruno/go-toyrobot/src/toyrobot/unit"
)

func TestNewTable(t *testing.T) {
	var board Tabler
	board = NewTable()
	if board == nil {
		t.Errorf("Could not instantiate a new table object")
	}
}

func TestReportNoPlace(t *testing.T) {
	// REPORT

	board := NewTable()
	output := board.Report()
	if output != msgRobotNotPlaced {
		t.Errorf("Expected robot to not be placed, but received [%s]", board.Report())
	}
}

func TestScenarioPlaceReport(t *testing.T) {
	// PLACE 0,0,NORTH
	// REPORT

	board := NewTable()
	board.Place(unit.X(1), unit.Y(2), compass.SOUTH)
	output := board.Report()
	if output != "1,2,SOUTH" {
		t.Errorf("Expected robot to be placed with location 0,0,NORTH but received [%s]", board.Report())
	}
}

func TestPlaceOutsideBoundsReport(t *testing.T) {
	// PLACE 5,5,SOUTH
	// REPORT

	board := NewTable()
	board.Place(unit.X(5), unit.Y(5), compass.SOUTH)
	output := board.Report()
	if output != msgRobotNotPlaced {
		t.Errorf("Expected robot to not be placed, but received [%s]", board.Report())
	}
}

func TestPlaceOneInboundsOneOutsideBoundsReport(t *testing.T) {
	// PLACE 4,5,SOUTH
	// PLACE 5,5,EAST
	// REPORT

	board := NewTable()
	board.Place(unit.X(4), unit.Y(5), compass.SOUTH)
	board.Place(unit.X(5), unit.Y(5), compass.EAST)
	output := board.Report()
	if output != msgRobotNotPlaced {
		t.Errorf("Expected robot to be placed with location 4,5,SOUTH but received [%s]", board.Report())
	}
}

func TestDoublePlaceValid(t *testing.T) {
	// PLACE 3,4,SOUTH
	// PLACE 2,3,EAST
	// REPORT

	board := NewTable()
	board.Place(unit.X(3), unit.Y(4), compass.SOUTH)
	board.Place(unit.X(2), unit.Y(3), compass.EAST)
	output := board.Report()
	if output != "2,3,EAST" {
		t.Errorf("Expected robot to be placed with location 2,3,EAST but received [%s]", board.Report())
	}
}

func TestMoveNotPlaced(t *testing.T) {
	board := NewTable()
	board.Move()
	output := board.Report()
	if output != msgRobotNotPlaced {
		t.Errorf("Expected robot to be not be placed but received [%s]", board.Report())
	}
}

func TestLeftNotPlaced(t *testing.T) {
	board := NewTable()
	board.Left()
	output := board.Report()
	if output != msgRobotNotPlaced {
		t.Errorf("Expected robot to be not be placed but received [%s]", board.Report())
	}
}

func TestRightNotPlaced(t *testing.T) {
	board := NewTable()
	board.Right()
	output := board.Report()
	if output != msgRobotNotPlaced {
		t.Errorf("Expected robot to be not be placed but received [%s]", board.Report())
	}
}

func TestPlaceEveryLocation(t *testing.T) {
	for f := compass.NORTH; f <= compass.WEST; f++ {
		for x := unit.X(0); x < unit.X(_tableLengthX); x = x + 1 {
			for y := unit.Y(0); y < unit.Y(_tableLengthX); y = y + 1 {
				board := NewTable()
				board.Place(x, y, f)
				location := board.Report()
				if location != fmt.Sprintf("%s,%s,%s", x, y, f) {
					t.Errorf("Place failed.  Expected location %s,%s,%s but received %s", x, y, f, location)
				}
			}
		}
	}
}

func TestMove(t *testing.T) {

	cases := []struct {
		placeX   unit.X
		placeY   unit.Y
		placeF   compass.Direction
		location string
	}{
		{
			placeX:   unit.X(3),
			placeY:   unit.Y(3),
			placeF:   compass.NORTH,
			location: "3,4,NORTH",
		},
		{
			placeX:   unit.X(3),
			placeY:   unit.Y(3),
			placeF:   compass.SOUTH,
			location: "3,2,SOUTH",
		},
		{
			placeX:   unit.X(3),
			placeY:   unit.Y(3),
			placeF:   compass.EAST,
			location: "4,3,EAST",
		},
		{
			placeX:   unit.X(3),
			placeY:   unit.Y(3),
			placeF:   compass.WEST,
			location: "2,3,WEST",
		},
	}

	for _, c := range cases {
		board := NewTable()
		board.Place(c.placeX, c.placeY, c.placeF)
		board.Move()
		location := board.Report()
		if location != c.location {
			t.Errorf("Move failed.  Expected location %s but received %s", c.location, location)
		}
	}

}

func TestMoveOutOfBounds(t *testing.T) {

	test := func(x unit.X, y unit.Y, f compass.Direction) {
		board := NewTable()
		board.Place(x, y, f)
		board.Move()
		location := board.Report()
		if location != fmt.Sprintf("%s,%s,%s", x, y, f) {
			t.Errorf("Move failed.  Expected location %s,%s,%s but received %s", x, y, f, location)
		}
	}

	var (
		x unit.X
		y unit.Y
	)

	for f := compass.NORTH; f <= compass.WEST; f++ {

		if f == compass.NORTH {
			y = unit.Y(_tableLengthY - 1)
			for x = unit.X(0); x < unit.X(_tableLengthX); x++ {
				test(x, y, f)
			}
		}

		if f == compass.EAST {
			x = unit.X(_tableLengthX - 1)
			for y = unit.Y(0); y < unit.Y(_tableLengthY); y++ {
				test(x, y, f)
			}
		}

		if f == compass.SOUTH {
			y = unit.Y(0)
			for x = unit.X(0); x < unit.X(_tableLengthX); x++ {
				test(x, y, f)
			}
		}

		if f == compass.WEST {
			x = unit.X(0)
			for y = unit.Y(0); y < unit.Y(_tableLengthY); y++ {
				test(x, y, f)
			}
		}

	}
}

func TestLeft(t *testing.T) {
	var output compass.Direction

	for f := compass.NORTH; f <= compass.WEST; f++ {
		board := NewTable()
		board.Place(3, 4, f)
		board.Left()
		location := board.Report()

		switch f {
		case compass.NORTH:
			output = compass.WEST
		case compass.EAST:
			output = compass.NORTH
		case compass.SOUTH:
			output = compass.EAST
		case compass.WEST:
			output = compass.SOUTH
		}

		if location != fmt.Sprintf("3,4,%s", output) {
			t.Errorf("Right failed.  Expected 3,4,%s but received %s", output, location)
		}
	}
}

func TestRight(t *testing.T) {

	var output compass.Direction

	for f := compass.NORTH; f <= compass.WEST; f++ {
		board := NewTable()
		board.Place(2, 3, f)
		board.Right()
		location := board.Report()

		switch f {
		case compass.NORTH:
			output = compass.EAST
		case compass.EAST:
			output = compass.SOUTH
		case compass.SOUTH:
			output = compass.WEST
		case compass.WEST:
			output = compass.NORTH
		}

		if location != fmt.Sprintf("2,3,%s", output) {
			t.Errorf("Right failed.  Expected 2,3,%s but received %s", output, location)
		}
	}
}
