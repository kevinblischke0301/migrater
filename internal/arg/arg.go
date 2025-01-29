package arg

import (
	"errors"
	"fmt"
	"strings"
)

type Command int8

const (
	MIGRATE Command = iota
)

func ParseArg(args []string) (Command, error) {
	n := len(args)

	if n == 1 {
		return MIGRATE, nil
	} else if n < 1 {
		return -1, errors.New("to few command-line arguments were given")
	} else if n > 2 {
		return -1, errors.New("to many command-line arguments were given")
	} else {
		return getCommand(args[1])
	}
}

func getCommand(arg string) (Command, error) {
	switch strings.ToLower(arg) {
	case "", "--migrate", "-migrate", "migrate":
		return MIGRATE, nil
	default:
		return -1, errors.New(fmt.Sprintf("the command-line argument \"%s\" isn't valid", arg))
	}
}
