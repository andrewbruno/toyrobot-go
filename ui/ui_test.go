package ui

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestDiplayIntro(t *testing.T) {
	var outOK bytes.Buffer
	displayIntro(&outOK)

	if welcome != outOK.String() {
		t.Errorf("Got %s\nbut expected\n%s", outOK.String(), welcome)
	}

}

func TestProgramQuits(t *testing.T) {

	expected := fmt.Sprintf("%sGoodbye\n", welcome)
	input := "\n\n\nq\n"

	var outOK bytes.Buffer

	ui := &UI{
		WriterOK:  &outOK,
		WriterErr: &outOK,
		ReaderIn:  strings.NewReader(input),
	}

	ui.Start()

	if outOK.String() != expected {
		t.Errorf("Expected\n%s but received\n%s", expected, outOK.String())
	}

}

func TestInvalidCommands(t *testing.T) {

	input := "lalala\n\n\nq\n"
	expected := fmt.Sprintf("%sOutput: Invalid input [lalala]\nGoodbye\n", welcome)

	var outOK bytes.Buffer

	ui := &UI{
		WriterOK:  &outOK,
		WriterErr: &outOK,
		ReaderIn:  strings.NewReader(input),
	}

	ui.Start()

	if outOK.String() != expected {
		t.Errorf("Expected\n%s but received\n%s", expected, outOK.String())
	}

}

func TestScenariosABCD(t *testing.T) {

	cases := []struct {
		input  string
		output string
	}{

		// ### Example A

		// PLACE 0,0,NORTH
		// MOVE
		// REPORT

		// Expected output: 0,1,NORTH
		{
			input:  "PLACE 0,0,NORTH\nMOVE\nREPORT\nq\n",
			output: fmt.Sprintf("%sOutput: 0,1,NORTH\nGoodbye\n", welcome), // Scenario A
		},

		// ### Example B

		// PLACE 0,0,NORTH
		// LEFT
		// REPORT

		// Expected output: 0,0,WEST
		{
			input:  "PLACE 0,0,NORTH\nLEFT\nREPORT\nq\n",
			output: fmt.Sprintf("%sOutput: 0,0,WEST\nGoodbye\n", welcome), // Scenario B
		},

		// ### Example C

		// PLACE 1,2,EAST
		// MOVE
		// MOVE
		// LEFT
		// MOVE
		// REPORT

		// Expected output:	3,3,NORTH

		{
			input:  "PLACE 1,2,EAST\nMOVE\nMOVE\nLEFT\nMOVE\nREPORT\nq\n",
			output: fmt.Sprintf("%sOutput: 3,3,NORTH\nGoodbye\n", welcome), // Scenario C
		},

		// ### Example D

		// PLACE 1,2,EAST
		// MOVE
		// MOVE
		// LEFT
		// MOVE
		// RIGHT
		// REPORT

		// Expected output:	3,3,EAST
		{
			input:  "PLACE 1,2,EAST\nMOVE\nMOVE\nLEFT\nMOVE\nRIGHT\nREPORT\nq\n", // Test RIGHT too
			output: fmt.Sprintf("%sOutput: 3,3,EAST\nGoodbye\n", welcome),
		},
	}

	for _, c := range cases {
		var outOK bytes.Buffer

		ui := &UI{
			WriterOK:  &outOK,
			WriterErr: &outOK,
			ReaderIn:  strings.NewReader(c.input),
		}

		ui.Start()

		if outOK.String() != c.output {
			t.Errorf("Expected\n%s\nbut received\n%s", c.output, outOK.String())
		}
	}
}
