package marsrover_test

import (
	"errors"
	"marsrover"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProvided(t *testing.T) {
	var (
		v        = testCases[0] // The provided test.
		reader   = strings.NewReader(v.Input)
		executor = marsrover.NewExecutor(reader)
		err      error
	)

	for err == nil {
		err = executor.Tick()
	}

	// Output locations for proof.
	for _, v := range executor.Rovers {
		t.Log(v.CurrentLocation)
	}

	if err == marsrover.ErrorEndOfInstructions {
		return
	}

	assert.Equal(t, v.ExpectedErr, unwrapError(err))
}

func Test(t *testing.T) {
	for _, v := range testCases {
		t.Run(v.Name, func(t *testing.T) {
			var (
				reader   = strings.NewReader(v.Input)
				executor = marsrover.NewExecutor(reader)
				err      error
			)

			for err == nil {
				err = executor.Tick()
			}

			if err == marsrover.ErrorEndOfInstructions {
				return
			}

			assert.Equal(t, v.ExpectedErr, unwrapError(err))
		})
	}
}

func unwrapError(err error) error {
	newErr := errors.Unwrap(err)
	for newErr != nil {
		err = newErr
		newErr = errors.Unwrap(newErr)
	}
	return err
}
