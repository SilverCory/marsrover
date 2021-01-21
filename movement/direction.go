package movement

import (
	"errors"
	"strings"
)

var (
	// ErrorInvalidDirection occurs when the direction is neither DirectionLeft or DirectionRight.
	ErrorInvalidDirection = errors.New("invalid direction")
)

// Direction is a direction to face.
type Direction int

// Left or Right
const (
	_ Direction = iota
	DirectionLeft
	DirectionRight
	// Deprecated DirectionFront is only for use with the interpreter and executor.
	DirectionFront
)

// Validate will ensure the direction is valid for use.
func (d Direction) Validate() error {
	switch d {
	case DirectionLeft, DirectionRight:
		return nil
	default:
		return ErrorInvalidDirection
	}
}

func DirectionFromString(s string) (Direction, error) {
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)

	switch s {
	case "L":
		return DirectionLeft, nil
	case "R":
		return DirectionRight, nil
	case "M":
		return DirectionFront, nil
	default:
		return 0, ErrorInvalidDirection
	}
}
