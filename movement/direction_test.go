package movement_test

import (
	. "marsrover/movement"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirection_Validate(t *testing.T) {
	type directionTestCase struct {
		Name      string
		Direction Direction
		Expected  error
	}

	var directionTestCases = []directionTestCase{
		{"Valid_DirectionLeft", DirectionLeft, nil},
		{"Valid_DirectionRight", DirectionRight, nil},
		{"Invalid_DirectionFront", DirectionFront, ErrorInvalidDirection},
		{"Invalid_Direction0", 0, ErrorInvalidDirection},
		{"Invalid_Direction5", 3, ErrorInvalidDirection},
		{"Invalid_Direction6", 4, ErrorInvalidDirection},
	}

	for _, v := range directionTestCases {
		t.Run(v.Name, func(t *testing.T) {
			assert.Equal(t, v.Expected, v.Direction.Validate())
		})
	}
}

func TestDirectionFromString(t *testing.T) {
	type directionTestCase struct {
		Name        string
		Input       string
		Expected    Direction
		ExpectedErr error
	}

	var directionTestCases = []directionTestCase{
		{"Valid_DirectionLeft", "L", DirectionLeft, nil},
		{"Valid_DirectionRight", "R", DirectionRight, nil},
		{"Valid_DirectionFront", "M", DirectionFront, nil},
		{"Valid_LowerDirectionLeft", "l", DirectionLeft, nil},
		{"Valid_LowerDirectionRight", "r", DirectionRight, nil},
		{"Valid_LowerDirectionFront", "m", DirectionFront, nil},
		{"Invalid_DirectionF", "F", 0, ErrorInvalidDirection},
		{"Invalid_Direction5", "5", 0, ErrorInvalidDirection},
		{"Invalid_Direction_", " ", 0, ErrorInvalidDirection},
	}

	for _, v := range directionTestCases {
		t.Run(v.Name, func(t *testing.T) {
			direct, err := DirectionFromString(v.Input)
			assert.Equal(t, v.Expected, direct)
			assert.Equal(t, v.ExpectedErr, err)
		})
	}
}
