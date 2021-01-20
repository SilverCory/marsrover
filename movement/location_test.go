package movement_test

import (
	. "marsrover/movement"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	locationEmpty    = Location{}
	locationNorth, _ = NewLocation(1, 2, OrientationNorth, 3, 3)
	locationEast, _  = NewLocation(1, 2, OrientationEast, 3, 3)
	locationSouth, _ = NewLocation(1, 2, OrientationSouth, 3, 3)
	locationWest, _  = NewLocation(1, 2, OrientationWest, 3, 3)

	locationOutOfBoundsXHigh, _     = NewLocation(2, 1, OrientationNorth, 1, 1)
	locationOutOfBoundsYHigh, _     = NewLocation(1, 2, OrientationNorth, 1, 1)
	locationOutOfBoundsXNegative, _ = NewLocation(-1, 0, OrientationNorth, 3, 3)
	locationOutOfBoundsYNegative, _ = NewLocation(0, -1, OrientationNorth, 3, 3)
)

func TestLocation_Turn(t *testing.T) {
	type locationTurnTestCase struct {
		Name          string
		Origin        Location
		Direction     Direction
		Expected      Location
		ExpectedError error
	}

	var locationTurnTestCases = []locationTurnTestCase{
		{"Valid_North_To_Left", locationNorth, DirectionLeft, locationWest, nil},
		{"Valid_North_To_Right", locationNorth, DirectionRight, locationEast, nil},
		{"Valid_South_To_Left", locationSouth, DirectionLeft, locationEast, nil},
		{"Valid_South_To_Right", locationSouth, DirectionRight, locationWest, nil},
		{"Valid_West_To_Left", locationWest, DirectionLeft, locationSouth, nil},
		{"Valid_West_To_Right", locationWest, DirectionRight, locationNorth, nil},
		{"Valid_East_To_Left", locationEast, DirectionLeft, locationNorth, nil},
		{"Valid_East_To_Right", locationEast, DirectionRight, locationSouth, nil},
		// Orientation to invalid direction (0)
		{"Invalid_North_To_0", locationEast, 0, locationEmpty, ErrorInvalidDirection},
		{"Invalid_South_To_0", locationSouth, 0, locationEmpty, ErrorInvalidDirection},
		{"Invalid_West_To_0", locationWest, 0, locationEmpty, ErrorInvalidDirection},
		{"Invalid_East_To_0", locationEast, 0, locationEmpty, ErrorInvalidDirection},
		// Orientation to invalid direction (3)
		{"Invalid_North_To_0", locationEast, 3, locationEmpty, ErrorInvalidDirection},
		{"Invalid_South_To_0", locationSouth, 3, locationEmpty, ErrorInvalidDirection},
		{"Invalid_West_To_0", locationWest, 3, locationEmpty, ErrorInvalidDirection},
		{"Invalid_East_To_0", locationEast, 3, locationEmpty, ErrorInvalidDirection},
		// Empty to direction
		{"Invalid_Empty_To_Right", locationEmpty, DirectionLeft, locationEmpty, ErrorInvalidOrientation},
		{"Invalid_Empty_To_Right", locationEmpty, DirectionRight, locationEmpty, ErrorInvalidOrientation},
	}

	for _, v := range locationTurnTestCases {
		t.Run(v.Name, func(t *testing.T) {
			res, err := v.Origin.Turn(v.Direction)
			assert.Equal(t, res, v.Expected)
			assert.Equal(t, err, v.ExpectedError)
		})
	}
}

func TestLocation_String(t *testing.T) {
	type locationStringTestCase struct {
		Name     string
		Location Location
		Expected string
	}

	var locationStringTestCases = []locationStringTestCase{
		{"Valid_LocationNorth", locationNorth, "1 2 N"},
		{"Valid_LocationEast", locationEast, "1 2 E"},
		{"Valid_LocationSouth", locationSouth, "1 2 S"},
		{"Valid_LocationWest", locationWest, "1 2 W"},
		{"Valid_LocationEmpty", locationEmpty, "0 0 Orientation{0}"},
	}

	for _, v := range locationStringTestCases {
		t.Run(v.Name, func(t *testing.T) {
			assert.Equal(t, v.Expected, v.Location.String())
		})
	}
}

func TestLocation_Move(t *testing.T) {
	type locationMoveTestCase struct {
		Name        string
		Location    Location
		ExpectedFn  func(Location) Location
		ExpectedErr error
	}

	var locationMoveTestCases = []locationMoveTestCase{
		{"Valid_LocationNorth", locationNorth, setRelative(0, 1), nil},
		{"Valid_LocationEast", locationEast, setRelative(1, 0), nil},
		{"Valid_LocationSouth", locationSouth, setRelative(0, -1), nil},
		{"Valid_LocationWest", locationWest, setRelative(-1, 0), nil},
	}

	for _, v := range locationMoveTestCases {
		t.Run(v.Name, func(t *testing.T) {
			var l, err = v.Location.Move()
			assert.Equal(t, v.ExpectedFn(v.Location), l)
			assert.Equal(t, v.ExpectedErr, err)
		})
	}
}

func TestLocation_Validate(t *testing.T) {
	type locationValidateTestCase struct {
		Name     string
		Location Location
		Expected error
	}

	var locationValidateTestCases = []locationValidateTestCase{
		{"Valid_LocationNorth", locationNorth, nil},
		{"Valid_LocationEast", locationEast, nil},
		{"Valid_LocationSouth", locationSouth, nil},
		{"Valid_LocationWest", locationWest, nil},
		{"Valid_LocationEmpty", locationEmpty, ErrorInvalidOrientation},
		{"Valid_Location_Bounds_XHigh", locationOutOfBoundsXHigh, ErrorLocationOutOfBounds},
		{"Valid_Location_Bounds_YHigh", locationOutOfBoundsYHigh, ErrorLocationOutOfBounds},
		{"Valid_Location_Bounds_XNegative", locationOutOfBoundsXNegative, ErrorLocationOutOfBounds},
		{"Valid_Location_Bounds_YNegative", locationOutOfBoundsYNegative, ErrorLocationOutOfBounds},
	}

	for _, v := range locationValidateTestCases {
		t.Run(v.Name, func(t *testing.T) {
			assert.Equal(t, v.Location.Validate(), v.Expected)
		})
	}
}

// doNothing returns a function that will return the same location supplied.
func doNothing() func(Location) Location {
	return func(l Location) Location {
		return l
	}
}

// setRelative returns a function that will set relative coordinates to the supplied Location.
func setRelative(X, Y int) func(Location) Location {
	return func(l Location) Location {
		l.X += X
		l.Y += Y
		return l
	}
}
