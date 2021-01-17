package movement

import "errors"

var (
	ErrorInvalidDirection = errors.New("invalid direction")
)

// Direction is a direction to face.
type Direction int

// Left or Right
const (
	_ Direction = iota
	DirectionLeft
	DirectionRight
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
