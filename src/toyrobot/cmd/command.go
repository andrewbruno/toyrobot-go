package cmd

import (
	"regexp"
)

const (
	cmdPlace           = "^PLACE (?P<x>[0-4]),(?P<y>[0-4]),(?P<f>NORTH|EAST|SOUTH|WEST)$"
	cmdPlaceMatchCount = 4 // input plus x, y, f

	cmdOther           = "^(?P<cmd>MOVE|LEFT|RIGHT|REPORT)$"
	cmdAtherMatchCount = 2 // input plus other cmd
)

var (
	regExpPlace = regexp.MustCompile(cmdPlace)
	regExpOther = regexp.MustCompile(cmdOther)
)

// ParseInput returns the action commands from an input string,
// as a array list of individual strings.
// Input must be valid, or regex will not match, and hence return nil
func ParseInput(input string) []string {

	commands := parsePlace(input)
	if commands != nil {
		return commands
	}

	return parseOther(input)
}

// parsePlace tests if the command is a valid PLACE command
// If invalid returns nil
// If valid returns a string array
//  e.g.  PLACE 0,1,SOUTH -> ["PLACE","0","1","SOUTH"]
func parsePlace(input string) []string {
	matchPlace := regExpPlace.FindStringSubmatch(input)

	if len(matchPlace) == cmdPlaceMatchCount {
		x := ""
		y := ""
		f := ""

		for i, name := range regExpPlace.SubexpNames() {
			if i > 0 && i <= len(matchPlace) {
				switch name {
				case "x":
					x = matchPlace[i]
				case "y":
					y = matchPlace[i]
				case "f":
					f = matchPlace[i]
				}
			}
		}
		if len(x) != 0 && len(y) != 0 && len(f) != 0 {
			return []string{"PLACE", x, y, f}
		}
	}

	return nil
}

// parseOther tests if the command is a valid command other than PLACE
// If invalid returns nil
// If valid returns a string array or one length, with first command in array [0]
//  e.g.  MOVE -> ["MOVE"]
func parseOther(input string) []string {
	match := regExpOther.FindStringSubmatch(input)

	if len(match) != cmdAtherMatchCount {
		return nil
	}

	var otherCmd []string

	for i, name := range regExpOther.SubexpNames() {
		if i > 0 && i <= len(match) {
			switch name {
			case "cmd":
				otherCmd = []string{match[i]}
			}
		}
	}

	return otherCmd
}
