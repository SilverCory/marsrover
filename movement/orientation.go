package movement

import (
	"errors"
	"fmt"
)

var (
	ErrorInvalidOrientation = errors.New("invalid orientation")
)

// Orientation is an ordinal direction.
type Orientation int

// The 4 main ordinal directions.
const (
	_ Orientation = iota
	OrientationNorth
	OrientationEast
	OrientationSouth
	OrientationWest
)

// Validate will ensure the orientation is ready for use.
func (o Orientation) Validate() error {
	switch o {
	case OrientationNorth, OrientationEast, OrientationSouth, OrientationWest:
		return nil
	default:
		return ErrorInvalidOrientation
	}
}

func (o Orientation) String() string {
	switch o {
	case OrientationNorth:
		return "N"
	case OrientationEast:
		return "E"
	case OrientationSouth:
		return "S"
	case OrientationWest:
		return "W"
	default:
		return fmt.Sprintf("Orientation{%d}", o)
	}
}
