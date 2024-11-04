package rover

import (
	"fmt"

	"github.com/MrTanguy/rover/planet"
	"github.com/MrTanguy/rover/utils"
)

type Rover interface {
	Forward()
	Backward()
	TurnRight()
	TurnLeft()
	ExecuteCommand(cmd string)
}

var directions = []string{"N", "E", "S", "W"}
var directionMap = map[string][2]int{
	"N": {0, 1},
	"S": {0, -1},
	"E": {1, 0},
	"W": {-1, 0},
}

type RoverImpl struct {
	X, Y        int
	Orientation string
	Planet      planet.Planet
}

func (r *RoverImpl) Forward() {
	change := directionMap[r.Orientation]
	r.X = r.Planet.WrapX(r.X + change[0])
	r.Y = r.Planet.WrapY(r.Y + change[1])
	fmt.Printf("Rover position: (%d, %d), Orientation: %s\n", r.X, r.Y, r.Orientation)
}

func (r *RoverImpl) Backward() {
	change := directionMap[r.Orientation]
	r.X = r.Planet.WrapX(r.X - change[0])
	r.Y = r.Planet.WrapY(r.Y - change[1])
	fmt.Printf("Rover position: (%d, %d), Orientation: %s\n", r.X, r.Y, r.Orientation)
}

func (r *RoverImpl) TurnRight() {
	idx := utils.FindIndex(directions, r.Orientation)
	idx = (idx + 1) % 4
	r.Orientation = directions[idx]
	fmt.Printf("Rover position: (%d, %d), Orientation: %s\n", r.X, r.Y, r.Orientation)
}

func (r *RoverImpl) TurnLeft() {
	idx := utils.FindIndex(directions, r.Orientation)
	idx = (idx + 3) % 4
	r.Orientation = directions[idx]
	fmt.Printf("Rover position: (%d, %d), Orientation: %s\n", r.X, r.Y, r.Orientation)
}

func (r *RoverImpl) ExecuteCommand(command string) {
	switch command {
	case "L":
		r.TurnLeft()
	case "R":
		r.TurnRight()
	case "F":
		r.Forward()
	case "B":
		r.Backward()
	}
}
