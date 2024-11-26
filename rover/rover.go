package rover

import (
	"fmt"

	planet "github.com/MrTanguy/rover/topologie"
	"github.com/MrTanguy/rover/utils"
)

type Rover interface {
	Forward() State
	Backward() State
	TurnRight() State
	TurnLeft() State
}

type State struct {
	X, Y        int
	Orientation string
}

func (s State) String() string {
	return fmt.Sprintf("rover is at x: %d y: %d orientation: %s", s.X, s.Y, s.Orientation)
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

func (r *RoverImpl) state() State {
	return State{X: r.X, Y: r.Y, Orientation: r.Orientation}
}

func (r *RoverImpl) Forward() State {
	change := directionMap[r.Orientation]
	r.X = r.Planet.WrapX(r.X + change[0])
	r.Y = r.Planet.WrapY(r.Y + change[1])
	return r.state()
}

func (r *RoverImpl) Backward() State {
	change := directionMap[r.Orientation]
	r.X = r.Planet.WrapX(r.X - change[0])
	r.Y = r.Planet.WrapY(r.Y - change[1])
	return r.state()
}

func (r *RoverImpl) TurnRight() State {
	idx := utils.FindIndex(directions, r.Orientation)
	idx = (idx + 1) % 4
	r.Orientation = directions[idx]
	return r.state()
}

func (r *RoverImpl) TurnLeft() State {
	idx := utils.FindIndex(directions, r.Orientation)
	idx = (idx + 3) % 4
	r.Orientation = directions[idx]
	return r.state()
}
