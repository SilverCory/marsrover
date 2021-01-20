package interpreter

import "errors"

var (
	// ErrorInvalidState occurs when an invalid state occurs.
	ErrorInvalidState = errors.New("invalid state")
	// ErrorUnexpectedState occurs when the current state isn't expected.
	ErrorUnexpectedState = errors.New("unexpected state")
)

type State int

const (
	_ State = iota
	StateSetPlateau
	StateCreateRover
	StateManipulateRover
)

// GetInstruction will return a new Instruction relevant to the current state.
//
// Returns ErrorInvalidState if this is called on an invalid state.
func (s State) GetInstruction() (Instruction, error) {
	switch s {
	case StateSetPlateau:
		return new(InstructionSetPlateau), nil
	case StateCreateRover:
		return new(InstructionCreateRover), nil
	case StateManipulateRover:
		return new(InstructionManipulateRover), nil
	default:
		return nil, ErrorInvalidState
	}
}

// GetNext will return the next State.
//
// Returns ErrorInvalidState if this is called on an invalid state.
func (s State) GetNext() (State, error) {
	switch s {
	case StateSetPlateau:
		return StateCreateRover, nil
	case StateCreateRover:
		return StateManipulateRover, nil
	case StateManipulateRover:
		return StateCreateRover, nil
	default:
		return 0, ErrorInvalidState
	}
}
