package interpreter

import (
	"bytes"
	"errors"
	"fmt"
	"marsrover/movement"
	"strings"
)

var (
	// ErrorInvalidManipulation occurs when an invalid instruction is supplied for manipulation.
	ErrorInvalidManipulation = errors.New("invalid manipulation")
)

// Used to ensure interface.
var _ Instruction = &InstructionManipulateRover{}

// InstructionManipulateRover is the instruction data for manipulating a rover.
type InstructionManipulateRover struct {
	Instruction
	Manipulations []movement.Direction
}

// Decode parses the supplied string into a InstructionManipulateRover.
func (i *InstructionManipulateRover) Decode(s string) (err error) {
	parts := strings.Split(s, " ")
	if len(parts) != 1 || strings.TrimSpace(parts[0]) == "" {
		return ErrorInstructionInvalidLength
	}

	byteParts := bytes.ToUpper([]byte(parts[0]))
	for _, v := range byteParts {
		manipulation, err := movement.DirectionFromString(string(v))
		if err != nil {
			// Allow for better userfeedback.
			if err == movement.ErrorInvalidDirection {
				err = ErrorInvalidManipulation
			}

			return fmt.Errorf("invalid data: %q: %w", v, err)
		}

		i.Manipulations = append(i.Manipulations, manipulation)
	}

	return nil
}
