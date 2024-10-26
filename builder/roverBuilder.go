package roverBuilder

import (
	"github.com/MrTanguy/rover/planet"
	"github.com/MrTanguy/rover/rover"
)

type MyRover struct {
	*rover.RoverImpl
}

func Build(rover_x int, rover_y int, rover_orientation string, planet_x int, planet_y int) *rover.RoverImpl {
	if rover_orientation == "" {
		rover_orientation = "N"
	}

	my_planet := &planet.Planet{Width: planet_x, Height: planet_y}
	my_rover := &rover.RoverImpl{X: rover_x, Y: rover_y, Orientation: rover_orientation, Planet: my_planet}

	return my_rover
}

func (r *MyRover) SetOrientation(new_orientation string) {
	r.Orientation = new_orientation
}
