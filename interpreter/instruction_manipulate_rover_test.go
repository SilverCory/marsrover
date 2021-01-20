package interpreter_test

import (
	"marsrover/movement"
	"testing"

	. "marsrover/interpreter"

	"github.com/stretchr/testify/assert"
)

func TestInstructionManipulateRover_Decode(t *testing.T) {
	type instructionManipulateRoverTestCase struct {
		Name        string
		Input       string
		Expected    []movement.Direction
		ExpectedErr error
	}

	var manipulateRoverTestCases = []instructionManipulateRoverTestCase{
		{"Valid_manipulate_LLLLL", "LLLLL", []movement.Direction{1, 1, 1, 1, 1}, nil},
		{"Valid_manipulate_RRRRR", "RRRRR", []movement.Direction{2, 2, 2, 2, 2}, nil},
		{"Valid_manipulate_RRRRR", "MMMMM", []movement.Direction{3, 3, 3, 3, 3}, nil},
		{"Valid_manipulate_LRFRLFFLR", "LRMRLMML", []movement.Direction{1, 2, 3, 2, 1, 3, 3, 1}, nil},
		{"Invalid_manipulate_PPPPP", "PPPPP", nil, ErrorInvalidManipulation},
		{"Invalid_manipulate_TooLong", "a a a", nil, ErrorInstructionInvalidLength},
		{"Invalid_manipulate_TooShort", "", nil, ErrorInstructionInvalidLength},
	}

	for _, v := range manipulateRoverTestCases {
		t.Run(v.Name, func(t *testing.T) {
			var instruction = new(InstructionManipulateRover)
			var err = instruction.Decode(v.Input)
			if err != nil {
				err = unwrapError(err)
			}
			assert.Equal(t, v.Expected, instruction.Manipulations)
			assert.Equal(t, v.ExpectedErr, err)
		})
	}
}
