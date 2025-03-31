package cmd_test

import (
	"testing"

	"github.com/andrewbruno/go-toyrobot/cmd"
)

func TestParseInput(t *testing.T) {

	cases := []struct {
		input  string
		output []string
	}{
		{
			input:  "PLACE 0,0,NORTH",
			output: []string{"PLACE", "0", "0", "NORTH"},
		},
		{
			input:  "PLACE 1,2,SOUTH",
			output: []string{"PLACE", "1", "2", "SOUTH"},
		},
		{
			input:  "PLACE 6,5,EAST",
			output: nil,
		},
		{
			input:  "PLACE 5,5,east",
			output: nil,
		},
		{
			input:  "PLACE 5, 5, EAST",
			output: nil,
		},
		{
			input:  "MOVE",
			output: []string{"MOVE"},
		},
		{
			input:  "LEFT",
			output: []string{"LEFT"},
		},
		{
			input:  "RIGHT",
			output: []string{"RIGHT"},
		},
		{
			input:  "REPORT",
			output: []string{"REPORT"},
		},
		{
			input:  "REPORT a",
			output: nil,
		},
		{
			input:  "PLACE 0a,0,NORTH",
			output: nil,
		},
	}

	for _, c := range cases {
		output := cmd.ParseInput(c.input)

		if output == nil && c.output != nil {
			t.Errorf("Unable to parse input [%s]", c.input)
			continue
		}

		if output != nil && c.output == nil {
			t.Errorf("Was able to parse input [%s] when it was expected to not be parseable", c.input)
			continue
		}

		if len(output) != len(c.output) {
			t.Errorf("Expected [%s] but received [%s] for input %s", c.output, output, c.input)
			continue
		}

		for i := range output {
			if output[i] != c.output[i] {
				t.Errorf("Expected %s but received %s for input [%s]", c.output, output, c.input)
				continue
			}
		}
	}
}
