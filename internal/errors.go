package internal

import (
	"fmt"
	"strings"
)

type UnknownInstructionErr struct {
	cmd byte
}

func NewUnknownInstructionErr(cmd byte) *UnknownInstructionErr {
	return &UnknownInstructionErr{cmd: cmd}
}

func (e *UnknownInstructionErr) Error() string {
	return "unknown instruction: " + string(e.cmd)
}

type UnevenParenthesesErr struct {
	openCnt, closeCnt int
}

func NewUnevenParenthesesErr(code string) *UnevenParenthesesErr {
	return &UnevenParenthesesErr{
		openCnt:  strings.Count(code, "["),
		closeCnt: strings.Count(code, "]"),
	}
}

func (e *UnevenParenthesesErr) Error() string {
	return fmt.Sprintf("uneven number of parentheses: %d opening and %d closing", e.openCnt, e.closeCnt)
}

type MemoryOverflowErr struct {
	memSize int
}

func NewMemoryOverflowErr() *MemoryOverflowErr {
	return &MemoryOverflowErr{
		memSize: len(mem),
	}
}

func (e MemoryOverflowErr) Error() string {
	return fmt.Sprintf("memory overflow (current mem is %d bytes long)", e.memSize)
}
