package marsrover

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"marsrover/interpreter"
	"marsrover/movement"
	"reflect"
	"strings"
)

var (
	// ErrorEndOfInstructions occurs when there are no more instructions to scan.
	ErrorEndOfInstructions = errors.New("no more instructions")
)

// Executor is the main executor that will handle instructions and execute them.
type Executor struct {
	Rovers []*Rover

	plateauMaxX int
	plateauMaxY int

	currentState interpreter.State
	scanner      *bufio.Scanner
}

// NewExecutor returns a new and ready instance of the executor using the supplied reader.
func NewExecutor(reader io.Reader) Executor {
	return Executor{
		currentState: interpreter.StateSetPlateau,
		scanner:      bufio.NewScanner(reader),
	}
}

// Tick will step forward one tick of time, read and execute an instruction.
func (e *Executor) Tick() error {
	if !e.scanner.Scan() {
		return ErrorEndOfInstructions
	}

	instruct, err := e.currentState.GetInstruction()
	if err != nil {
		return fmt.Errorf("state GetInstruction: %w", err)
	}

	instructionStr := e.scanner.Text()
	err = e.scanner.Err()
	if err != nil {
		return fmt.Errorf("scanner err: %w", err)
	}

	instructionStr = strings.TrimSpace(instructionStr)
	if instructionStr == "" {
		return nil
	}

	err = instruct.Decode(instructionStr)
	if err != nil {
		return fmt.Errorf("instruction (%q) Decode: %w", reflect.TypeOf(instruct).Name(), err)
	}

	return e.executeInstruction(instruct)
}

func (e *Executor) executeInstruction(i interpreter.Instruction) error {
	switch i.(type) {
	case *interpreter.InstructionSetPlateau:
		return e.executeInstructionSetPlateau(i.(*interpreter.InstructionSetPlateau))
	case *interpreter.InstructionCreateRover:
		return e.executeInstructionCreateRover(i.(*interpreter.InstructionCreateRover))
	case *interpreter.InstructionManipulateRover:
		return e.executeInstructionManipulateRover(i.(*interpreter.InstructionManipulateRover))
	default:
		panic(fmt.Sprintf("invalid instruction: %v", i))
	}
}

func (e *Executor) executeInstructionSetPlateau(i *interpreter.InstructionSetPlateau) error {
	e.plateauMaxX = i.MaxX
	e.plateauMaxY = i.MaxY

	return e.nextState()
}

func (e *Executor) executeInstructionCreateRover(i *interpreter.InstructionCreateRover) error {
	location, err := movement.NewLocation(i.X, i.Y, i.Orientation, e.plateauMaxX, e.plateauMaxY)
	if err != nil {
		return fmt.Errorf("executeInstructionCreateRover: %w", err)
	}

	e.Rovers = append(e.Rovers, NewRover(location))
	return e.nextState()
}

func (e *Executor) executeInstructionManipulateRover(i *interpreter.InstructionManipulateRover) error {
	rover := e.getWorkingRover()
	if rover == nil {
		panic(fmt.Errorf("no working rover for manipulate command"))
	}

	for _, v := range i.Manipulations {
		var newLocation movement.Location
		var err error

		switch v {
		case movement.DirectionLeft, movement.DirectionRight:
			newLocation, err = rover.CurrentLocation.Turn(v)
		case movement.DirectionFront:
			newLocation, err = rover.CurrentLocation.Move()
		}

		if err != nil {
			// See Readme.md#Assumptions.
			return err
		}

		rover.LocationHistory = append(rover.LocationHistory, rover.CurrentLocation)
		rover.CurrentLocation = newLocation
	}

	return e.nextState()
}

func (e *Executor) nextState() error {
	nextState, err := e.currentState.GetNext()
	if err != nil {
		return fmt.Errorf("nextState: %w", err)
	}

	e.currentState = nextState
	return nil
}

func (e *Executor) getWorkingRover() *Rover {
	if len(e.Rovers) < 1 {
		return nil
	}

	return e.Rovers[len(e.Rovers)-1]
}
