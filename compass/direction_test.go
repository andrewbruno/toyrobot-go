package compass

import "testing"
import "fmt"

func TestEdgeLeft(t *testing.T) {
	newDirection := Left(NORTH)
	if newDirection != WEST {
		t.Errorf("Expected WEST but got %s", newDirection)
	}
}

func TestEdgeRight(t *testing.T) {
	newDirection := Right(WEST)
	if newDirection != NORTH {
		t.Errorf("Expected NORTH but got %s", newDirection)
	}
}

func Test180Right(t *testing.T) {
	newDirection := Right(NORTH)
	newDirection = Right(newDirection)
	if newDirection != SOUTH {
		t.Errorf("Expected SOUTH but got %s", newDirection)
	}
}

func Test180Left(t *testing.T) {
	newDirection := Left(NORTH)
	newDirection = Left(newDirection)
	if newDirection != SOUTH {
		t.Errorf("Expected SOUTH but got %s", newDirection)
	}
}

func Test450Left(t *testing.T) {
	newDirection := Left(NORTH)
	newDirection = Left(newDirection)
	newDirection = Left(newDirection)
	newDirection = Left(newDirection)
	newDirection = Left(newDirection)
	if newDirection != WEST {
		t.Errorf("Expected EAST but got %s", newDirection)
	}
}

func Test450Right(t *testing.T) {
	newDirection := Right(NORTH)
	newDirection = Right(newDirection)
	newDirection = Right(newDirection)
	newDirection = Right(newDirection)
	newDirection = Right(newDirection)
	if newDirection != EAST {
		t.Errorf("Expected WEST but got %s", newDirection)
	}
}

func TestStringOutput(t *testing.T) {

	cases := []struct {
		input    Direction
		expected string
	}{
		{
			input:    NORTH,
			expected: "NORTH",
		},
		{
			input:    EAST,
			expected: "EAST",
		},
		{
			input:    SOUTH,
			expected: "SOUTH",
		},
		{
			input:    WEST,
			expected: "WEST",
		},
	}

	for _, c := range cases {

		output := fmt.Sprintf("%s", c.input)

		if c.expected != output {
			t.Errorf("Location String failed; got %s, expected %s", output, c.expected)
		}
	}
}

func TestParseDirection(t *testing.T) {
	cases := []struct {
		input    string
		expected Direction
		err      error
	}{
		{
			input:    "NORTH",
			expected: NORTH,
		},
		{
			input:    "EAST",
			expected: EAST,
		},
		{
			input:    "SOUTH",
			expected: SOUTH,
		},
		{
			input:    "WEST",
			expected: WEST,
		},
		{
			input: "WESTwe",
			err:   fmt.Errorf("Can not parse direction [%s]", "WESTwe"),
		},
	}

	for _, c := range cases {
		output, err := ParseDirection(c.input)

		if err != nil {
			if err.Error() != c.err.Error() {
				t.Errorf("Expected error message [%s], but got [%s]", c.err.Error(), err.Error())
			}
		} else {
			if c.expected != *output {
				t.Errorf("Expected [%s], but got [%s]", c.expected, output)
			}
		}
	}
}
