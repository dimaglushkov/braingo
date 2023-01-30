package internal

import (
	"fmt"
	"strings"
)

var commands = map[byte]func(string) error{
	'r': InterResetMem,
	'm': InterSetIOMode,
	'd': InterDumpMem,
}

func ExecInteractive(cmd string) error {
	call := cmd[1]
	if _, ok := commands[call]; !ok {
		return NewUnknownInstructionErr(call, true)
	}

	return commands[call](cmd)
}

func InterResetMem(cmd string) error {
	args := strings.Fields(cmd)
	if len(args) == 1 {
		ResetMem(len(mem))
		return nil
	}
	memSize, err := atoi(args[1])
	if err != nil {
		return err
	}
	ResetMem(memSize)
	return nil
}

func InterSetIOMode(cmd string) error {
	return nil
}

func InterDumpMem(cmd string) error {
	args := strings.Fields(cmd)
	if len(args) != 3 {
		return new(WrongNumberOfArgsErr)
	}
	si, err := atoi(args[1])
	if err != nil {
		return err
	}
	ei, err := atoi(args[2])
	if err != nil {
		return err
	}

	if si > ei {
		return nil
	}

	for i := si; i < ei; i++ {
		fmt.Printf("%v.", mem[i])
	}
	fmt.Println()

	return nil
}
