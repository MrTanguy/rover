package main

import "github.com/MrTanguy/rover/rover"

func main() {
	// Initialize a rover at position (2, 2) facing North on a 5x5 toroidal planet
	var myRover rover.Rover = &rover.RoverImpl{X: 2, Y: 2, Orientation: "N", PlanetWidth: 5, PlanetHeight: 5}

	commands := []string{"M", "R", "M", "M", "L", "M", "B"}

	// Execute commands and print the state after each command
	for _, cmd := range commands {
		myRover.ExecuteCommand(cmd)
		myRover.PrintState()
	}
}
