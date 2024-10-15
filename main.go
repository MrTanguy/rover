package main

import "fmt"

func main() {

}

func (r *RoverImpl) Turn(direction string, findIndex func([]string, string) int) {
	idx := findIndex(directions, r.Orientation)

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

func (r *RoverImpl) ExecuteCommand(command string, findIndex func([]string, string) int) {
	switch command {
	case "L":
		r.Turn("L", findIndex)
	case "R":
		r.Turn("R", findIndex)
	case "M":
		r.Move(1)
	case "B":
		r.Move(-1)
	}
}

func (r *RoverImpl) PrintState() {
	fmt.Printf("Rover position: (%d, %d), Orientation: %s\n", r.X, r.Y, r.Orientation)
}
