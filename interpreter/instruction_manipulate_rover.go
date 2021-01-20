package interpreter

import (
	"MarsRover/movement"
	"bytes"
	"errors"
	"fmt"
	"strings"
)

var (
	// ErrorInvalidManipulation occurs when an invalid instruction is supplied for manipulation.
	ErrorInvalidManipulation = errors.New("invalid manipulation")
)

// Used to ensure interface.
var _ Instruction = &InstructionManipulateRover{}

type InstructionManipulateRover struct {
	Instruction
	Manipulations []movement.Direction
}

func (i *InstructionManipulateRover) Decode(s string) (err error) {
	parts := strings.Split(s, " ")
	if len(parts) != 1 || strings.TrimSpace(parts[0]) == "" {
		return ErrorInstructionInvalidLength
	}

	byteParts := bytes.ToUpper([]byte(parts[0]))
	for _, v := range byteParts {
		manipulation, err := manipulationFromByte(v)
		if err != nil {
			return fmt.Errorf("invalid data: %q: %w", v, err)
		}

		i.Manipulations = append(i.Manipulations, manipulation)
	}

	return nil
}

// TODO move this to where it belongs.
func manipulationFromByte(b byte) (movement.Direction, error) {
	switch b {
	case 'L':
		return movement.DirectionLeft, nil
	case 'R':
		return movement.DirectionRight, nil
	case 'F':
		return movement.DirectionFront, nil
	default:
		return 0, ErrorInvalidManipulation
	}
}
