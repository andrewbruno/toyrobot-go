package unit_test

import (
	"fmt"
	"testing"
	"toyrobot/unit"
)

func TestUnitXString(t *testing.T) {
	input := unit.X(2)
	expect := "2"
	output := fmt.Sprintf("%s", input)

	if expect != output {
		t.Errorf("Expected %s but got %s", expect, output)
	}
}

func TestUnitYString(t *testing.T) {
	input := unit.Y(255)
	expect := "255"
	output := fmt.Sprintf("%s", input)

	if expect != output {
		t.Errorf("Expected %s but got %s", expect, output)
	}
}
