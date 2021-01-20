package interpreter_test

import (
	"errors"
	. "marsrover/interpreter"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstructionSetPlateau_Decode(t *testing.T) {
	type instructionSetPlateauTestCase struct {
		Name        string
		Input       string
		X, Y        int
		ExpectedErr error
	}

	var setPlateauTestCases = []instructionSetPlateauTestCase{
		{"Valid_SetPlateau_1_1", "1 1", 1, 1, nil},
		{"Valid_SetPlateau_5_5", "5 5", 5, 5, nil},
		{"Invalid_SetPlateau_a_b", "a b", 0, 0, strconv.ErrSyntax},
		{"Invalid_SetPlateau_1_b", "1 b", 1, 0, strconv.ErrSyntax},
		{"Invalid_SetPlateau_", " ", 0, 0, strconv.ErrSyntax},
		{"Invalid_SetPlateau_TooLong", "a a a", 0, 0, ErrorInstructionInvalidLength},
		{"Invalid_SetPlateau_TooShort", "a", 0, 0, ErrorInstructionInvalidLength},
	}

	for _, v := range setPlateauTestCases {
		t.Run(v.Name, func(t *testing.T) {
			var instruction = new(InstructionSetPlateau)
			var err = instruction.Decode(v.Input)
			if err != nil {
				err = unwrapError(err)
			}

			assert.Equal(t, v.X, instruction.MaxX)
			assert.Equal(t, v.Y, instruction.MaxY)
			assert.Equal(t, v.ExpectedErr, err)
		})
	}
}

func unwrapError(err error) error {
	newErr := errors.Unwrap(err)
	for newErr != nil {
		err = newErr
		newErr = errors.Unwrap(newErr)
	}
	return err
}
