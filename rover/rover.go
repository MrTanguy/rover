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
var directionMap = map[string][2]int{
	"N": {0, -1},
	"S": {0, 1},
	"E": {1, 0},
	"W": {-1, 0},
}

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
	change := directionMap[r.Orientation]
	r.X = (r.X + change[0]*step + r.PlanetWidth) % r.PlanetWidth
	r.Y = (r.Y + change[1]*step + r.PlanetHeight) % r.PlanetHeight
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
