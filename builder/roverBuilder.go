package roverBuilder

import (
	"github.com/MrTanguy/rover/planet"
	"github.com/MrTanguy/rover/rover"
)

type MyRover struct {
	*rover.RoverImpl
}

func Build(opts ...RoverOption) *rover.RoverImpl {
	const (
		defaultX           = 0
		defaultY           = 0
		defaultOrientation = "N"
		defaultPlanetX     = 0
		defaultPlanetY     = 0
	)

	my_planet := &planet.Planet{Width: defaultPlanetX, Height: defaultPlanetY}

	my_rover := &rover.RoverImpl{X: defaultX, Y: defaultY, Orientation: defaultOrientation, Planet: my_planet}

	for _, opt := range opts {
		opt(my_rover)
	}

	return my_rover
}
