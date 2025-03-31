package unit_test

import (
	"testing"

	"github.com/andrewbruno/go-toyrobot/unit"
)

func TestUnitXString(t *testing.T) {
	input := unit.X(2)
	expect := "2"
	output := input.String()

	if expect != output {
		t.Errorf("Expected %s but got %s", expect, output)
	}
}

func TestUnitYString(t *testing.T) {
	input := unit.Y(255)
	expect := "255"
	output := input.String()

	if expect != output {
		t.Errorf("Expected %s but got %s", expect, output)
	}
}
