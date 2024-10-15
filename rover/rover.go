package rover

import (
	"fmt"

	"github.com/MrTanguy/rover/utils"
)

type Rover interface {
	Move(step int)
	Turn(direction string)
	ExecuteCommand(cmd string)
	PrintState()
}

var directions = []string{"N", "E", "S", "W"}

// RoverImpl is the implementation of the Rover interface
type RoverImpl struct {
	X, Y         int
	Orientation  string
	PlanetWidth  int
	PlanetHeight int
}

func (r *RoverImpl) Turn(direction string) {
	idx := utils.FindIndex(directions, r.Orientation)
	if direction == "L" {
		idx = (idx + 3) % 4
	} else if direction == "R" {
		idx = (idx + 1) % 4
	}

	r.Orientation = directions[idx]
}

func (r *RoverImpl) Move(step int) {
	switch r.Orientation {
	case "N":
		r.Y = (r.Y - step + r.PlanetHeight) % r.PlanetHeight
	case "S":
		r.Y = (r.Y + step) % r.PlanetHeight
	case "E":
		r.X = (r.X + step) % r.PlanetWidth
	case "W":
		r.X = (r.X - step + r.PlanetWidth) % r.PlanetWidth
	}
}

func (r *RoverImpl) ExecuteCommand(command string) {
	switch command {
	case "L":
		r.Turn("L")
	case "R":
		r.Turn("R")
	case "M":
		r.Move(1)
	case "B":
		r.Move(-1)
	}
}

func (r *RoverImpl) PrintState() {
	fmt.Printf("Rover position: (%d, %d), Orientation: %s\n", r.X, r.Y, r.Orientation)
}
