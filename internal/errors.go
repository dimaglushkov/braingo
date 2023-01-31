package internal

import (
	"errors"
	"fmt"
)

var (
	errWrongNumberOfArguments    = errors.New("not enough arguments to performing this action")
	errMemorySizeExceeded        = errors.New("memory size exceeded")
	errInvalidMemorySize         = errors.New("invalid memory size value")
	errUnknownInteractiveCommand = errors.New("unknown interactive command")
	errUnknownIOMode             = errors.New("unknown IO format")
)

type UnknownInstructionError error

func NewUnknownInstructionError(cmd byte) UnknownInstructionError {
	return UnknownInstructionError(fmt.Errorf("unknown instruction: %c", cmd))
}
