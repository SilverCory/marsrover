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
