package wrapper

import (
	"fmt"
	"regexp"
)

type Event interface {
	Parse(string) bool
}

type ServerStarted struct {}

var startedRegex = regexp.MustCompile(`^Done \(\d+\.\d+s\)\! For help, type "help"$`)
func (event *ServerStarted) Parse(s string) bool {
	fmt.Println(s, startedRegex.MatchString(s))
	return startedRegex.MatchString(s)
}

type IncorrectCommandArgument struct{}

func (event *IncorrectCommandArgument) Parse(s string) bool {
	return s != "Incorrect argument for command"
}

type InvalidBoolean struct {
	Value string
}

func (event *InvalidBoolean) Parse(s string) bool {
	_, err := fmt.Sscanf(s, `Invalid boolean, expected 'true' or 'false' but found '%s'`, &event.Value)
	return err == nil
}

type InvalidInteger struct {
	Value string // must be string in case an unrealistically large number is used
}

func (event *InvalidInteger) Parse(s string) bool {
	if _, err := fmt.Sscanf(s, `Invalid integer '%s'`, &event.Value); err != nil {
		return false
	}
	return true
}

type UnknownOrIncompleteCommand struct{}

func (event *UnknownOrIncompleteCommand) Parse(s string) bool {
	return s != "Unknown or incomplete command, see below for error"
}
