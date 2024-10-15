package rover_test

import (
	"testing"

	"github.com/MrTanguy/rover/rover"
	"github.com/stretchr/testify/assert"
)

func TestRoverTurnLeft(t *testing.T) {
	// Initialize the rover at position (0, 0) facing North
	myRover := &rover.RoverImpl{X: 0, Y: 0, Orientation: "N", PlanetWidth: 5, PlanetHeight: 5}

	// Turn left
	myRover.Turn("L")

	// Check if the orientation is West after turning left
	assert.Equal(t, "W", myRover.Orientation)
}

func TestRoverTurnRight(t *testing.T) {
	// Initialize the rover at position (0, 0) facing North
	myRover := &rover.RoverImpl{X: 0, Y: 0, Orientation: "N", PlanetWidth: 5, PlanetHeight: 5}

	// Turn right
	myRover.Turn("R")

	// Check if the orientation is East after turning right
	assert.Equal(t, "E", myRover.Orientation)
}

func TestRoverMoveForward(t *testing.T) {
	// Initialize the rover at position (2, 2) facing North
	myRover := &rover.RoverImpl{X: 2, Y: 2, Orientation: "N", PlanetWidth: 5, PlanetHeight: 5}

	// Move forward
	myRover.Move(1)

	// Check if the position is (2, 1) after moving forward
	assert.Equal(t, 2, myRover.X)
	assert.Equal(t, 1, myRover.Y)
}

func TestRoverMoveBackward(t *testing.T) {
	// Initialize the rover at position (2, 2) facing North
	myRover := &rover.RoverImpl{X: 2, Y: 2, Orientation: "N", PlanetWidth: 5, PlanetHeight: 5}

	// Move backward
	myRover.Move(-1)

	// Check if the position is (2, 3) after moving backward
	assert.Equal(t, 2, myRover.X)
	assert.Equal(t, 3, myRover.Y)
}

func TestRoverExecuteCommand(t *testing.T) {
	// Initialize the rover at position (0, 0) facing North
	myRover := &rover.RoverImpl{X: 0, Y: 0, Orientation: "N", PlanetWidth: 5, PlanetHeight: 5}

	// Execute a series of commands: Move, Turn Right, Move
	myRover.ExecuteCommand("M")
	myRover.ExecuteCommand("R")
	myRover.ExecuteCommand("M")

	// The rover should be at position (1, 4) facing East
	assert.Equal(t, 1, myRover.X)
	assert.Equal(t, 4, myRover.Y)
	assert.Equal(t, "E", myRover.Orientation)
}

func TestRoverToroidalWrapping(t *testing.T) {
	// Initialize the rover at position (0, 0) facing North on a 5x5 grid
	myRover := &rover.RoverImpl{X: 0, Y: 0, Orientation: "N", PlanetWidth: 5, PlanetHeight: 5}

	// Move forward, it should wrap around the planet
	myRover.Move(1)

	// The rover should now be at position (0, 4) due to toroidal wrapping
	assert.Equal(t, 0, myRover.X)
	assert.Equal(t, 4, myRover.Y)
}
