package interpreter_test

import (
	. "marsrover/interpreter"
	"marsrover/movement"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstructionCreateRover_Decode(t *testing.T) {
	type instructionCreateRoverTestCase struct {
		Name        string
		Input       string
		X, Y        int
		Orientation movement.Orientation
		ExpectedErr error
	}

	var createRoverTestCases = []instructionCreateRoverTestCase{
		{"Valid_CreateRover_1_1_N", "1 1 N", 1, 1, 1, nil},
		{"Valid_CreateRover_1_1_E", "1 1 E", 1, 1, 2, nil},
		{"Valid_CreateRover_1_1_S", "1 1 S", 1, 1, 3, nil},
		{"Valid_CreateRover_1_1_W", "1 1 W", 1, 1, 4, nil},
		{"Valid_CreateRover_5_5_n", "5 5 n", 5, 5, 1, nil},
		{"Valid_CreateRover_5_5_e", "5 5 e", 5, 5, 2, nil},
		{"Valid_CreateRover_5_5_s", "5 5 s", 5, 5, 3, nil},
		{"Valid_CreateRover_5_5_w", "5 5 w", 5, 5, 4, nil},
		{"Invalid_CreateRover_a_b_1", "a b 1", 0, 0, 0, strconv.ErrSyntax},
		{"Invalid_CreateRover_1_b_3", "1 b 3", 1, 0, 0, strconv.ErrSyntax},
		{"Invalid_CreateRover_1_1_F", "1 1 F", 1, 1, 0, movement.ErrorInvalidOrientation},
		{"Invalid_CreateRover_1_1_O", "1 1 O", 1, 1, 0, movement.ErrorInvalidOrientation},
		{"Invalid_CreateRover_", " ", 0, 0, 0, ErrorInstructionInvalidLength},
		{"Invalid_CreateRover_TooLong", "a a A a", 0, 0, 0, ErrorInstructionInvalidLength},
		{"Invalid_CreateRover_TooShort", "a", 0, 0, 0, ErrorInstructionInvalidLength},
	}

	for _, v := range createRoverTestCases {
		t.Run(v.Name, func(t *testing.T) {
			var instruction = new(InstructionCreateRover)
			var err = instruction.Decode(v.Input)
			if err != nil {
				err = unwrapError(err)
			}

			assert.Equal(t, v.X, instruction.X)
			assert.Equal(t, v.Y, instruction.Y)
			assert.Equal(t, v.Orientation, instruction.Orientation)
			assert.Equal(t, v.ExpectedErr, err)
		})
	}
}
