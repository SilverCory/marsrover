package interpreter

import "errors"

var (
	// ErrorInstructionInvalidLength occurs when too many parts are supplied to an instruction.
	ErrorInstructionInvalidLength = errors.New("instruction data was an invalid length")
)

// Instruction is an instruction
type Instruction interface {
	// Decode will decode the instruction
	Decode(string) error
}
