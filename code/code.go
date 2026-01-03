package code

import (
	"encoding/binary"
	"fmt"
)

type Opcode byte

const (
	OpConstant Opcode = iota
)

type Instructions []byte

func (ins Instructions) String() string {
	return ""
}

type Definition struct {
	Name         string
	OperandWidth []int
}

var definitions = map[Opcode]*Definition{
	OpConstant: {"OpConstant", []int{2}},
}

func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefinied", op)
	}

	return def, nil
}

func Make(op Opcode, operands ...int) []byte {
	def, ok := definitions[op]
	if !ok {
		return []byte{}
	}

	instructionLen := 1
	for _, w := range def.OperandWidth {
		instructionLen += w
	}

	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)

	offset := 1
	for i, operand := range operands {
		width := def.OperandWidth[i]
		switch width {
		case 2:
			// todo: implement my own big-endian conversion
			binary.BigEndian.PutUint16(instruction[offset:], uint16(operand))
		}
		offset += width
	}

	return instruction
}
