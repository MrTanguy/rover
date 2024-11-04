package roverBuilder

import (
	"github.com/MrTanguy/rover/planet"
	"github.com/MrTanguy/rover/rover"
)

type RoverOption func(r *rover.RoverImpl)

func WithOrientation(Orientation string) RoverOption {
	return func(r *rover.RoverImpl) {
		r.Orientation = Orientation
	}
}

func WithPosition(X int, Y int) RoverOption {
	return func(r *rover.RoverImpl) {
		r.X = X
		r.Y = Y
	}
}

func WithPlanet(Height int, Width int) RoverOption {
	return func(r *rover.RoverImpl) {
		r.Planet, _ = planet.New(planet.WithSize(Height, Width))
	}
}

func WithInfinityPlanet() RoverOption {
	return func(r *rover.RoverImpl) {
		r.Planet, _ = planet.New(planet.WithInfity())
	}
}
