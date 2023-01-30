package internal

import (
	"fmt"
	"strings"
)

type UnknownInstructionErr struct {
	cmd   byte
	inter bool
}

func NewUnknownInstructionErr(cmd byte, inter bool) *UnknownInstructionErr {
	return &UnknownInstructionErr{
		cmd:   cmd,
		inter: inter,
	}
}

func (e *UnknownInstructionErr) Error() string {
	if !e.inter {
		return "unknown instruction: " + string(e.cmd)
	} else {
		return "unknown interactive instruction: \\" + string(e.cmd)
	}
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

func (e *MemoryOverflowErr) Error() string {
	return fmt.Sprintf("memory overflow (current mem is %d bytes long)", e.memSize)
}

type WrongNumberOfArgsErr struct{}

func (e *WrongNumberOfArgsErr) Error() string {
	return fmt.Sprintf("not enough arguments to performing this action")
}
