package rover_test

import (
	"testing"

	roverBuilder "github.com/MrTanguy/rover/builder"
	"github.com/stretchr/testify/assert"
)

func TestRoverTurnLeft(t *testing.T) {

	// Rover orienté Nord en 0;0
	// Sur une planète de 5;5
	myRover := roverBuilder.Build(roverBuilder.WithPlanet(5, 5))

	myRover.Turn("L")

	assert.Equal(t, "W", myRover.Orientation)
}

func TestRoverTurnRight(t *testing.T) {

	// Rover orienté Nord en 0;0
	// Sur une planète de 5;5
	myRover := roverBuilder.Build(roverBuilder.WithPlanet(5, 5))

	myRover.Turn("R")
	assert.Equal(t, "E", myRover.Orientation)
}

func TestRoverMoveForward(t *testing.T) {

	// Rover orienté Nord en 0;0
	// Sur une planète de 5;5
	myRover := roverBuilder.Build(roverBuilder.WithPlanet(5, 5))

	myRover.Move(1)
	assert.Equal(t, 0, myRover.X)
	assert.Equal(t, 1, myRover.Y)
}

func TestRoverMoveBackward(t *testing.T) {

	// Rover orienté Nord en 0;0
	// Sur une planète de 5;5
	myRover := roverBuilder.Build(roverBuilder.WithPlanet(5, 5))

	myRover.Move(-1)

	assert.Equal(t, 0, myRover.X)
	assert.Equal(t, 4, myRover.Y)
}

func TestRoverExecuteCommand(t *testing.T) {

	// Rover orienté Nord en 0;0
	// Sur une planète de 5;5
	myRover := roverBuilder.Build(roverBuilder.WithPlanet(5, 5))

	myRover.ExecuteCommand("M")
	myRover.ExecuteCommand("R")
	myRover.ExecuteCommand("M")

	assert.Equal(t, 1, myRover.X)
	assert.Equal(t, 1, myRover.Y)
	assert.Equal(t, "E", myRover.Orientation)
}

func TestRoverToroidalWrapping(t *testing.T) {

	// Rover orienté Nord en 0;0
	// Sur une planète de 5;5
	myRover := roverBuilder.Build(roverBuilder.WithPlanet(5, 5))

	myRover.Move(1)

	assert.Equal(t, 0, myRover.X)
	assert.Equal(t, 1, myRover.Y)
}

func TestRoverWrappingOnPlanetEdges(t *testing.T) {

	// Rover orienté Est en 4;4
	// Sur une planète de 5;5
	myRover := roverBuilder.Build(roverBuilder.WithPosition(4, 4), roverBuilder.WithOrientation("E"), roverBuilder.WithPlanet(5, 5))

	myRover.Move(1)
	assert.Equal(t, 0, myRover.X)
	assert.Equal(t, 4, myRover.Y)

	myRover.Turn("R")
	myRover.Move(1)
	assert.Equal(t, 0, myRover.X)
	assert.Equal(t, 3, myRover.Y)
}

func TestRoverTurnAndMove(t *testing.T) {

	// Rover orienté Nord en 2;2
	// Sur une planète de 5;5
	myRover := roverBuilder.Build(roverBuilder.WithPosition(2, 2), roverBuilder.WithPlanet(5, 5))

	myRover.Turn("R")
	myRover.Move(1)
	assert.Equal(t, 3, myRover.X)
	assert.Equal(t, 2, myRover.Y)

	myRover.Turn("L")
	myRover.Move(1)
	assert.Equal(t, 3, myRover.X)
	assert.Equal(t, 3, myRover.Y)
}
