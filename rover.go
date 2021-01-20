package MarsRover

import "MarsRover/movement"

// Rover is the location, and location history of the rover.
type Rover struct {
	CurrentLocation movement.Location
	LocationHistory []movement.Location
}

// NewRover will create a new instance of *Rover with the supplied location.
func NewRover(l movement.Location) *Rover {
	return &Rover{CurrentLocation: l}
}
