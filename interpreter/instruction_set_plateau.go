package interpreter

import (
	"fmt"
	"strconv"
	"strings"
)

// Used to ensure interface.
var _ Instruction = &InstructionSetPlateau{}

// InstructionSetPlateau is the instruction data for setting up the plateau.
type InstructionSetPlateau struct {
	Instruction
	MaxX int
	MaxY int
}

// Decode parses the supplied string into a InstructionSetPlateau.
func (i *InstructionSetPlateau) Decode(s string) (err error) {
	parts := strings.Split(s, " ")
	if len(parts) != 2 {
		return ErrorInstructionInvalidLength
	}

	if i.MaxX, err = strconv.Atoi(parts[0]); err != nil {
		return fmt.Errorf("invalid maxX: %w", err)
	}

	if i.MaxY, err = strconv.Atoi(parts[1]); err != nil {
		return fmt.Errorf("invalid maxY: %w", err)
	}

	return nil
}
