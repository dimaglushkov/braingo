package internal

import (
	"fmt"
	"strings"
)

const (
	DefaultMemSize = 30_000
)

var (
	mem          []byte
	ioMode       IOMode = IOModeDigit
	ptr                 = 0
	spacingChars        = [...]string{
		" ",
		"\n",
		"\t",
	}
)

func init() {
	ResetMem(DefaultMemSize)

}

func ResetMem(memSize int) int {
	if memSize <= 0 {
		memSize = DefaultMemSize
	}
	mem = make([]byte, memSize)
	return len(mem)
}

func SetIOMode(iom IOMode) {
	ioMode = iom
}

func Execute(cmd string) error {
	for _, sch := range spacingChars {
		cmd = strings.ReplaceAll(cmd, sch, "")
	}

	i := 0
	for i < len(cmd) {
		switch cmd[i] {
		case '>':
			ptr++
		case '<':
			ptr--
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
			return NewUnknownInstructionErr(cmd[i])
		}

		i++
	}
	return nil
}
