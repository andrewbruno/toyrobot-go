package main

import (
	"bytes"
	"testing"
)

func TestRun(t *testing.T) {
	// Mock input and output streams
	mockInput := bytes.NewBufferString("PLACE 0,0,NORTH\nMOVE\nREPORT\nQ\n")
	mockOutput := &bytes.Buffer{}
	mockError := &bytes.Buffer{}

	// Call the run function with mocked streams
	run(mockOutput, mockError, mockInput)

	// Validate the output
	output := mockOutput.String()
	if output == "" {
		t.Errorf("Expected output, but got empty string")
	}

	// Validate no errors
	errorOutput := mockError.String()
	if errorOutput != "" {
		t.Errorf("Expected no errors, but got: %s", errorOutput)
	}
}
