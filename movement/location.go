package movement

import (
	"errors"
	"fmt"
)

var (
	// ErrorLocationOutOfBounds occurs when location is invalid and is outside of it's defined perimeter.
	ErrorLocationOutOfBounds = errors.New("location out of bounds")
)

// Location holds data on position and orientation
type Location struct {
	X           int
	Y           int
	Orientation Orientation

	boundX int
	boundY int
}

// NewLocation will create a new location with the specified location, orientation and bounds.
//
// An error will be returned if any of the parameters are invalid or out of bounds.
func NewLocation(X int, Y int, orientation Orientation, boundX int, boundY int) (Location, error) {
	var ret = Location{
		X:           X,
		Y:           Y,
		Orientation: orientation,
		boundX:      boundX,
		boundY:      boundY,
	}
	return ret, ret.Validate()
}

// Turn will invoke an orientation change in the given direction and return a new Location.
//
// An error is returned if an invalid origin orientation is used.
// An error is returned if an invalid turn direction is supplied.
func (l Location) Turn(d Direction) (Location, error) {
	if err := d.Validate(); err != nil {
		return Location{}, err
	}

	// Ensure we are in a good place!
	if err := l.Orientation.Validate(); err != nil {
		return Location{}, err
	}

	switch d {
	case DirectionLeft:
		l.Orientation--
	case DirectionRight:
		l.Orientation++
	}

	// 0 is invalid so wrap around and head west.
	if l.Orientation <= 0 {
		l.Orientation = OrientationWest
	}
	// 5 is invalid so wrap around and head north.
	if l.Orientation >= 5 {
		l.Orientation = OrientationNorth
	}

	// Validate again for sanity sake.
	return l, d.Validate()
}

// Move will move one step in the direction of the orientation and return a new Location.
//
// An error is returned if the movement goes out of bounds.
func (l Location) Move() (Location, error) {
	switch l.Orientation {
	case OrientationNorth:
		l.Y++
	case OrientationSouth:
		l.Y--
	case OrientationEast:
		l.X++
	case OrientationWest:
		l.X--
	}

	return l, l.Validate()
}

func (l Location) String() string {
	return fmt.Sprintf("%d %d %s", l.X, l.Y, l.Orientation.String())
}

// Validate will ensure the location is ready for use and inside bounds.
func (l Location) Validate() error {
	if err := l.Orientation.Validate(); err != nil {
		return err
	}

	// Ensure within bounds.
	if l.X < 0 || l.Y < 0 || l.X > l.boundX || l.Y > l.boundY {
		return ErrorLocationOutOfBounds
	}

	return nil
}
