package movement_test

import (
	. "MarsRover/movement"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrientation_Validate(t *testing.T) {
	type orientationTestCase struct {
		Name        string
		Orientation Orientation
		Expected    error
	}

	var orientationTestCases = []orientationTestCase{
		{"Valid_OrientationNorth", OrientationNorth, nil},
		{"Valid_OrientationEast", OrientationEast, nil},
		{"Valid_OrientationSouth", OrientationSouth, nil},
		{"Valid_OrientationWest", OrientationWest, nil},
		{"Valid_Orientation0", 0, ErrorInvalidOrientation},
		{"Valid_Orientation5", 5, ErrorInvalidOrientation},
		{"Valid_Orientation6", 6, ErrorInvalidOrientation},
	}

	for _, v := range orientationTestCases {
		t.Run(v.Name, func(t *testing.T) {
			assert.Equal(t, v.Expected, v.Orientation.Validate())
		})
	}
}

func TestOrientation_String(t *testing.T) {
	type orientationTestCase struct {
		Name        string
		Orientation Orientation
		Expected    string
	}

	var orientationTestCases = []orientationTestCase{
		{"Valid_OrientationNorth", OrientationNorth, "N"},
		{"Valid_OrientationEast", OrientationEast, "E"},
		{"Valid_OrientationSouth", OrientationSouth, "S"},
		{"Valid_OrientationWest", OrientationWest, "W"},
		{"Valid_Orientation0", 0, "Orientation{0}"},
		{"Valid_Orientation5", 5, "Orientation{5}"},
	}

	for _, v := range orientationTestCases {
		t.Run(v.Name, func(t *testing.T) {
			assert.Equal(t, v.Expected, v.Orientation.String())
		})
	}
}
