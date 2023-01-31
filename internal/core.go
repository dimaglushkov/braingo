package internal

import (
	"fmt"
	"os"
)

const DefaultMemSize = 30_000

var (
	mem           = make([]byte, DefaultMemSize)
	ioMode IOMode = IOModeDigit
	ptr           = 0
)

func ExecuteFile(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	return Execute(string(content))
}

func Execute(cmd string) error {
	cmd = stripSpaces(cmd)

	i := 0
	for i < len(cmd) {
		switch cmd[i] {
		case '>':
			ptr++
			if ptr >= len(mem) {
				ptr = 0
			}
		case '<':
			ptr--
			if ptr < 0 {
				ptr = len(mem) - 1
			}
		case '+':
			mem[ptr]++
		case '-':
			mem[ptr]--
		case '.':
			fmt.Printf(string(ioMode), mem[ptr])
		case ',':
			if _, err := fmt.Scanf(string(ioMode), &mem[ptr]); err != nil {
				return err
			}
		case '[':
			if mem[ptr] == 0 {
				var cycle int
				for cycle, i = 1, i+1; cycle != 0; i++ {
					if cmd[i] == '[' {
						cycle++
					}
					if cmd[i] == ']' {
						cycle--
					}
				}
			}
		case ']':
			if mem[ptr] != 0 {
				var cycle int
				for cycle, i = -1, i-1; cycle != 0; i-- {
					if cmd[i] == '[' {
						cycle++
					}
					if cmd[i] == ']' {
						cycle--
					}
				}
			}
		default:
			return NewUnknownInstructionError(cmd[i])
		}

		i++
	}
	return nil
}
