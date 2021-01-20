package MarsRover_test

import (
	"MarsRover"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	for _, v := range testCases {
		t.Run(v.Name, func(t *testing.T) {
			var (
				reader   = strings.NewReader(v.Input)
				executor = MarsRover.NewExecutor(reader)
				err      error
			)

			for err == nil {
				err = executor.Tick()
			}

			if err == MarsRover.ErrorEndOfInstructions {
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
