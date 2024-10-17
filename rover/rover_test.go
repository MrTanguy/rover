package rover_test

import (
	"testing"

	"github.com/MrTanguy/rover/planet"
	"github.com/MrTanguy/rover/rover"
	"github.com/stretchr/testify/assert"
)

func TestRoverTurnLeft(t *testing.T) {
	myPlanet := planet.NewPlanet(5, 5)
	myRover := &rover.RoverImpl{X: 0, Y: 0, Orientation: "N", Planet: myPlanet}

	myRover.Turn("L")
	assert.Equal(t, "W", myRover.Orientation)
}

func TestRoverTurnRight(t *testing.T) {
	myPlanet := planet.NewPlanet(5, 5)
	myRover := &rover.RoverImpl{X: 0, Y: 0, Orientation: "N", Planet: myPlanet}

	myRover.Turn("R")
	assert.Equal(t, "E", myRover.Orientation)
}

func TestRoverMoveForward(t *testing.T) {
	myPlanet := planet.NewPlanet(5, 5)
	myRover := &rover.RoverImpl{X: 2, Y: 2, Orientation: "N", Planet: myPlanet}

	myRover.Move(1)
	assert.Equal(t, 2, myRover.X)
	assert.Equal(t, 1, myRover.Y)
}

func TestRoverMoveBackward(t *testing.T) {
	myPlanet := planet.NewPlanet(5, 5)
	myRover := &rover.RoverImpl{X: 2, Y: 2, Orientation: "N", Planet: myPlanet}

	myRover.Move(-1)
	assert.Equal(t, 2, myRover.X)
	assert.Equal(t, 3, myRover.Y)
}

func TestRoverExecuteCommand(t *testing.T) {
	myPlanet := planet.NewPlanet(5, 5)
	myRover := &rover.RoverImpl{X: 0, Y: 0, Orientation: "N", Planet: myPlanet}

	myRover.ExecuteCommand("M")
	myRover.ExecuteCommand("R")
	myRover.ExecuteCommand("M")

	assert.Equal(t, 1, myRover.X)
	assert.Equal(t, 4, myRover.Y)
	assert.Equal(t, "E", myRover.Orientation)
}

func TestRoverToroidalWrapping(t *testing.T) {
	myPlanet := planet.NewPlanet(5, 5)
	myRover := &rover.RoverImpl{X: 0, Y: 0, Orientation: "N", Planet: myPlanet}

	myRover.Move(1)
	assert.Equal(t, 0, myRover.X)
	assert.Equal(t, 4, myRover.Y)
}

func TestRoverWrappingOnPlanetEdges(t *testing.T) {
	myPlanet := planet.NewPlanet(5, 5)
	myRover := &rover.RoverImpl{X: 4, Y: 4, Orientation: "E", Planet: myPlanet}

	myRover.Move(1)
	assert.Equal(t, 0, myRover.X)
	assert.Equal(t, 4, myRover.Y)

	myRover.Turn("R")
	myRover.Move(1)
	assert.Equal(t, 0, myRover.X)
	assert.Equal(t, 0, myRover.Y)
}

func TestRoverTurnAndMove(t *testing.T) {
	myPlanet := planet.NewPlanet(5, 5)
	myRover := &rover.RoverImpl{X: 2, Y: 2, Orientation: "N", Planet: myPlanet}

	myRover.Turn("R")
	myRover.Move(1)
	assert.Equal(t, 3, myRover.X)
	assert.Equal(t, 2, myRover.Y)

	myRover.Turn("L")
	myRover.Move(1)
	assert.Equal(t, 3, myRover.X)
	assert.Equal(t, 1, myRover.Y)
}
