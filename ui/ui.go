package ui

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/andrewbruno/go-toyrobot/cmd"
	"github.com/andrewbruno/go-toyrobot/compass"
	"github.com/andrewbruno/go-toyrobot/table"
	"github.com/andrewbruno/go-toyrobot/unit"
)

// UI represents the CLI
type UI struct {
	ReaderIn  io.Reader
	WriterOK  io.Writer
	WriterErr io.Writer
}

const welcome = "\n-- Toy Robot Simulator -- \n\nPlease enter your commands:\n"

// A hashmap of robot moveable actions mapped by text command
var robotActions = map[string]func(*table.Table){
	"MOVE":  (*table.Table).Move,
	"LEFT":  (*table.Table).Left,
	"RIGHT": (*table.Table).Right,
}

// Start the CLI for the app
func (u *UI) Start() {

	displayIntro(u.WriterOK)

	board := table.NewTable()

	reader := bufio.NewReader(u.ReaderIn)

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		if input != "" {
			if strings.ToLower(input) == "q" {
				break
			}

			commands := cmd.ParseInput(input)

			if commands == nil {
				fmt.Fprintf(u.WriterErr, "Output: Invalid input [%s]\n", input)
				continue
			}

			u.doCommands(board, commands)

		}
		fmt.Fprint(u.WriterOK, "")
	} // for

	fmt.Fprint(u.WriterOK, "Goodbye\n")

}

// displayIntro displays the welcome message.
func displayIntro(writerOK io.Writer) {
	fmt.Fprint(writerOK, welcome)
}

// doCommands takes in an array list of string of commands
// e.g. [PLACE,0,0,NORTH]
func (u *UI) doCommands(board *table.Table, commands []string) {
	if commands[0] == "PLACE" && len(commands) == 4 {
		xStr, _ := strconv.ParseUint(commands[1], 10, 8)
		yStr, _ := strconv.ParseUint(commands[2], 10, 8)

		x := unit.X(xStr)
		y := unit.Y(yStr)

		f, _ := compass.ParseDirection(commands[3])

		board.Place(x, y, *f)
	} else if commands[0] == "REPORT" {
		fmt.Fprintf(u.WriterOK, "Output: %s\n", board.Report())
	} else {
		robotActions[commands[0]](board)
	}
}
