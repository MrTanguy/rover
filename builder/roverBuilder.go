package roverBuilder

import (
	"github.com/MrTanguy/rover/rover"
	planet "github.com/MrTanguy/rover/topologie"
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

	my_planet, err := planet.New(false, planet.WithSize(defaultPlanetX, defaultPlanetY))
	if err != nil {
		panic(err)
	}

	my_rover := &rover.RoverImpl{X: defaultX, Y: defaultY, Orientation: defaultOrientation, Planet: my_planet}

	for _, opt := range opts {
		opt(my_rover)
	}

	return my_rover
}
