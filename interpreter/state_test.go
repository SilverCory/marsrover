package interpreter_test

import (
	. "MarsRover/interpreter"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestState_GetInstruction(t *testing.T) {
	type stateInstructionTestCase struct {
		Name        string
		State       State
		Expected    reflect.Type
		ExpectedErr error
	}

	var stateInstructionTestCases = []stateInstructionTestCase{
		{"Valid_StateSetPlateau", StateSetPlateau, reflect.TypeOf(&InstructionSetPlateau{}), nil},
		{"Valid_StateCreateRover", StateCreateRover, reflect.TypeOf(&InstructionCreateRover{}), nil},
		{"Valid_StateManipulateRover", StateManipulateRover, reflect.TypeOf(&InstructionManipulateRover{}), nil},
		{"Invalid_State0", 0, nil, ErrorInvalidState},
		{"Invalid_State4", 4, nil, ErrorInvalidState},
	}

	for _, v := range stateInstructionTestCases {
		t.Run(v.Name, func(t *testing.T) {
			instruction, err := v.State.GetInstruction()
			assert.Equal(t, v.Expected, reflect.TypeOf(instruction))
			assert.Equal(t, v.ExpectedErr, err)
		})
	}
}

func TestState_GetNext(t *testing.T) {
	type stateNextTestCase struct {
		Name        string
		State       State
		Expected    State
		ExpectedErr error
	}

	var stateNextTestCases = []stateNextTestCase{
		{"Valid_StateSetPlateau", StateSetPlateau, StateCreateRover, nil},
		{"Valid_StateCreateRover", StateCreateRover, StateManipulateRover, nil},
		{"Valid_StateManipulateRover", StateManipulateRover, StateCreateRover, nil},
		{"Invalid_State0", 0, 0, ErrorInvalidState},
		{"Invalid_State4", 4, 0, ErrorInvalidState},
	}

	for _, v := range stateNextTestCases {
		t.Run(v.Name, func(t *testing.T) {
			nextState, err := v.State.GetNext()
			assert.Equal(t, v.Expected, nextState)
			assert.Equal(t, v.ExpectedErr, err)
		})
	}
}
