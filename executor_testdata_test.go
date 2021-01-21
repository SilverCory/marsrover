package marsrover_test

import (
	"marsrover/interpreter"
	"strconv"
)

const (
	DataValid1 = `5 5
1 2 N
LMLMLMLMM
3 3 E
MMRMMRMRRM
`
	DataValid2 = `15 15
7 7 N
MMMLLMMMRRMMMRRMMM`
	DataInvalidNoPlateau = `0 0
0 0 N
RR`
	DataInvalidPlateauFormat = `ffff
0 0 N
RR`
	DataInvalidRoverCreateNone = `15 15
ff ff 0
RR`
	DataInvalidRoverCreateFormat = `15 15
a
RR`
	DataInvalidManipulationFormat = `15 15
0 0 N
pp`
)

var testCases = []struct {
	Name        string
	Input       string
	ExpectedErr error
}{
	{
		Name:        "Provided Test (Valid 1)",
		Input:       DataValid1,
		ExpectedErr: nil,
	}, {
		Name:        "valid2",
		Input:       DataValid2,
		ExpectedErr: nil,
	}, {
		Name:        "DataInvalidPlateauFormat",
		Input:       DataInvalidPlateauFormat,
		ExpectedErr: interpreter.ErrorInstructionInvalidLength,
	}, {
		Name:        "DataInvalidNoPlateau",
		Input:       DataInvalidNoPlateau,
		ExpectedErr: interpreter.ErrorInvalidManipulation,
	}, {
		Name:        "DataInvalidRoverCreateFormat",
		Input:       DataInvalidRoverCreateFormat,
		ExpectedErr: interpreter.ErrorInstructionInvalidLength,
	}, {
		Name:        "DataInvalidRoverCreateNone",
		Input:       DataInvalidRoverCreateNone,
		ExpectedErr: strconv.ErrSyntax,
	}, {
		Name:        "DataInvalidManipulationFormat",
		Input:       DataInvalidManipulationFormat,
		ExpectedErr: interpreter.ErrorInvalidManipulation,
	},
}
