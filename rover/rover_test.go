package rover_test

import (
	"testing"

	roverBuilder "github.com/MrTanguy/rover/builder"
	"github.com/stretchr/testify/assert"
)

func TestRoverTurnLeftOnInfinityPlanet(t *testing.T) {

	// Rover orienté Nord en 0;0
	// Sur une planète infinie
	myRover := roverBuilder.Build(roverBuilder.WithInfinityPlanet())

	myRover.TurnLeft()

	assert.Equal(t, 0, myRover.X)
	assert.Equal(t, 0, myRover.Y)
	assert.Equal(t, "W", myRover.Orientation)
}

func TestRoverTurnRightOnInfinityPlanet(t *testing.T) {

	// Rover orienté Nord en 0;0
	// Sur une planète infinie
	myRover := roverBuilder.Build(roverBuilder.WithInfinityPlanet())

	myRover.TurnRight()

	assert.Equal(t, 0, myRover.X)
	assert.Equal(t, 0, myRover.Y)
	assert.Equal(t, "E", myRover.Orientation)
}

func TestRoverMoveForwardOnInfinityPlanet(t *testing.T) {

	// Rover orienté Nord en 0;0
	// Sur une planète infinie
	myRover := roverBuilder.Build(roverBuilder.WithInfinityPlanet())

	myRover.Forward()

	assert.Equal(t, 0, myRover.X)
	assert.Equal(t, 1, myRover.Y)
	assert.Equal(t, "N", myRover.Orientation)
}

func TestRoverMoveBackwardOnInfinityPlanet(t *testing.T) {

	// Rover orienté Nord en 0;0
	// Sur une planète infinie
	myRover := roverBuilder.Build(roverBuilder.WithInfinityPlanet())

	myRover.Backward()

	assert.Equal(t, 0, myRover.X)
	assert.Equal(t, -1, myRover.Y)
	assert.Equal(t, "N", myRover.Orientation)
}

func TestRoverMultipleCommandsOnInfinityPlanet(t *testing.T) {

	// Rover orienté Nord en 0;0
	// Sur une planète infinie
	myRover := roverBuilder.Build(roverBuilder.WithInfinityPlanet())

	myRover.Forward()
	myRover.Forward()
	myRover.Forward()
	myRover.TurnRight()
	myRover.Backward()
	myRover.Backward()
	myRover.TurnLeft()
	myRover.Forward()

	assert.Equal(t, -2, myRover.X)
	assert.Equal(t, 4, myRover.Y)
	assert.Equal(t, "N", myRover.Orientation)
}

func TestRoverTurnLeftOnToroidalPlanet(t *testing.T) {

	// Rover orienté Nord en 0;0
	// Sur une planète de 5;5
	myRover := roverBuilder.Build(roverBuilder.WithPlanet(5, 5))

	myRover.TurnLeft()

	assert.Equal(t, 0, myRover.X)
	assert.Equal(t, 0, myRover.Y)
	assert.Equal(t, "W", myRover.Orientation)
}

func TestRoverTurnRightOnToroidalPlanet(t *testing.T) {

	// Rover orienté Nord en 0;0
	// Sur une planète de 5;5
	myRover := roverBuilder.Build(roverBuilder.WithPlanet(5, 5))

	myRover.TurnRight()

	assert.Equal(t, 0, myRover.X)
	assert.Equal(t, 0, myRover.Y)
	assert.Equal(t, "E", myRover.Orientation)
}

func TestRoverMoveForwardOnToroidalPlanet(t *testing.T) {

	// Rover orienté Nord en 0;0
	// Sur une planète de 5;5
	myRover := roverBuilder.Build(roverBuilder.WithPlanet(5, 5))

	myRover.Forward()

	assert.Equal(t, 0, myRover.X)
	assert.Equal(t, 1, myRover.Y)
	assert.Equal(t, "N", myRover.Orientation)
}

func TestRoverMoveBackwardOnToroidalPlanet(t *testing.T) {

	// Rover orienté Nord en 0;0
	// Sur une planète de 5;5
	myRover := roverBuilder.Build(roverBuilder.WithPlanet(5, 5))

	myRover.Backward()

	assert.Equal(t, 0, myRover.X)
	assert.Equal(t, 4, myRover.Y)
	assert.Equal(t, "N", myRover.Orientation)
}

func TestRoverWrappingOnToroidalPlanetEdges(t *testing.T) {

	// Rover orienté Est en 4;4
	// Sur une planète de 5;5
	myRover := roverBuilder.Build(roverBuilder.WithPosition(4, 4), roverBuilder.WithOrientation("E"), roverBuilder.WithPlanet(5, 5))

	myRover.Forward()
	assert.Equal(t, 0, myRover.X)
	assert.Equal(t, 4, myRover.Y)
	assert.Equal(t, "E", myRover.Orientation)

	myRover.TurnLeft()
	myRover.Forward()
	assert.Equal(t, 0, myRover.X)
	assert.Equal(t, 0, myRover.Y)
	assert.Equal(t, "N", myRover.Orientation)
}

func TestRoverMultipleCommandsOnToroidalPlanet(t *testing.T) {

	// Rover orienté Nord en 2;2
	// Sur une planète de 3;3
	myRover := roverBuilder.Build(roverBuilder.WithPosition(2, 2), roverBuilder.WithPlanet(3, 3))

	myRover.TurnRight()
	myRover.Forward()
	assert.Equal(t, 0, myRover.X)
	assert.Equal(t, 2, myRover.Y)
	assert.Equal(t, "E", myRover.Orientation)

	myRover.TurnRight()
	myRover.Forward()
	assert.Equal(t, 0, myRover.X)
	assert.Equal(t, 1, myRover.Y)
	assert.Equal(t, "S", myRover.Orientation)
}
