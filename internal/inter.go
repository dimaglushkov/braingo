package internal

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

const helpMessage = `-----
Braingo - brainfuck interpreter written in go
Available interactive commands:
\f <file>  - run code from <file>;
\r <size>  - reset pointer & memory and set its size to <size>;
\m <d/c>   - change current IO format: 
				d - print values as digits, 
				c - print ascii symbol represented by value;
\d <f> <t> - print memory values at cells with indices from <f> to <t>;
\h         - print this message;
-----`

func ExecuteInteractive() error {
	fmt.Println(helpMessage)
	in := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")

		x, err := in.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
		}
		if x[0] == '\\' {
			if err = execInterCommand(x); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = Execute(x); err != nil {
				fmt.Println(err)
			}
		}
	}
}

func execInterCommand(cmd string) error {
	args := strings.Fields(cmd)
	instr := args[0]

	switch instr {
	case "\\f":
		if len(args) != 2 {
			return errWrongNumberOfArguments
		}
		return ExecuteFile(args[1])
	case "\\r":
		return resetMem(args...)
	case "\\m":
		return setIOMode(args...)
	case "\\d":
		return dumpMem(args...)
	case "\\h":
		fmt.Println(helpMessage)
		return nil
	default:
		return errUnknownInteractiveCommand
	}
}

func resetMem(args ...string) error {
	ptr = 0
	if len(args) == 1 {
		for i := range mem {
			mem[i] = 0
		}
		return nil
	}
	memSize, err := atoi(args[1])
	if err != nil {
		return err
	}
	if memSize <= 0 {
		return errInvalidMemorySize
	}
	mem = make([]byte, memSize)
	return nil
}

func setIOMode(args ...string) error {
	if len(args) != 2 {
		return errWrongNumberOfArguments
	}
	switch m := args[1]; m {
	case "d":
		ioMode = IOModeDigit
	case "c":
		ioMode = IOModeChar
	default:
		return errUnknownIOMode
	}

	return nil
}

func dumpMem(args ...string) error {
	if len(args) != 3 {
		return errWrongNumberOfArguments
	}
	si, err := atoi(args[1])
	if err != nil {
		return err
	}
	ei, err := atoi(args[2])
	if err != nil {
		return err
	}

	if si > ei || si > len(mem) || ei > len(mem) || si < 0 || ei <= 0 {
		return errMemorySizeExceeded
	}

	for i := si; i < ei; i++ {
		fmt.Printf("%v ", mem[i])
	}
	fmt.Println()
	return nil
}
