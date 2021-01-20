package interpreter

import (
	"MarsRover/movement"
	"fmt"
	"strconv"
	"strings"
)

// Used to ensure interface.
var _ Instruction = &InstructionCreateRover{}

type InstructionCreateRover struct {
	Instruction
	X           int
	Y           int
	Orientation movement.Orientation
}

func (i *InstructionCreateRover) Decode(s string) (err error) {
	parts := strings.Split(s, " ")
	if len(parts) != 3 {
		return ErrorInstructionInvalidLength
	}

	if i.X, err = strconv.Atoi(parts[0]); err != nil {
		return fmt.Errorf("invalid X: %w", err)
	}

	if i.Y, err = strconv.Atoi(parts[1]); err != nil {
		return fmt.Errorf("invalid Y: %w", err)
	}

	if i.Orientation, err = movement.OrientationFromString(parts[2]); err != nil {
		return fmt.Errorf("invalid orientation: %w", err)
	}

	return nil
}
