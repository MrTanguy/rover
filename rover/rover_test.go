package rover_test

import (
	"testing"

	"github.com/MrTanguy/rover/rover"
	"github.com/stretchr/testify/assert"
)

func TestRoverTurnLeft(t *testing.T) {
	myRover := &rover.RoverImpl{X: 0, Y: 0, Orientation: "N", PlanetWidth: 5, PlanetHeight: 5}

	myRover.Turn("L")

	assert.Equal(t, "W", myRover.Orientation)
}

func TestRoverTurnRight(t *testing.T) {
	myRover := &rover.RoverImpl{X: 0, Y: 0, Orientation: "N", PlanetWidth: 5, PlanetHeight: 5}

	myRover.Turn("R")

	assert.Equal(t, "E", myRover.Orientation)
}

func TestRoverMoveForward(t *testing.T) {
	myRover := &rover.RoverImpl{X: 2, Y: 2, Orientation: "N", PlanetWidth: 5, PlanetHeight: 5}

	myRover.Move(1)

	assert.Equal(t, 2, myRover.X)
	assert.Equal(t, 1, myRover.Y)
}

func TestRoverMoveBackward(t *testing.T) {
	myRover := &rover.RoverImpl{X: 2, Y: 2, Orientation: "N", PlanetWidth: 5, PlanetHeight: 5}

	myRover.Move(-1)

	assert.Equal(t, 2, myRover.X)
	assert.Equal(t, 3, myRover.Y)
}

func TestRoverExecuteCommand(t *testing.T) {
	myRover := &rover.RoverImpl{X: 0, Y: 0, Orientation: "N", PlanetWidth: 5, PlanetHeight: 5}

	myRover.ExecuteCommand("M")
	myRover.ExecuteCommand("R")
	myRover.ExecuteCommand("M")

	assert.Equal(t, 1, myRover.X)
	assert.Equal(t, 4, myRover.Y)
	assert.Equal(t, "E", myRover.Orientation)
}

func TestRoverToroidalWrapping(t *testing.T) {
	myRover := &rover.RoverImpl{X: 0, Y: 0, Orientation: "N", PlanetWidth: 5, PlanetHeight: 5}

	myRover.Move(1)

	assert.Equal(t, 0, myRover.X)
	assert.Equal(t, 4, myRover.Y)
}
