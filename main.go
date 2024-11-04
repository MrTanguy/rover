package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MrTanguy/rover/planet"
	"github.com/MrTanguy/rover/rover"
)

func main() {

	planet := planet.New(planet.WithSize(5, 5))
	var myRover rover.Rover = &rover.RoverImpl{X: 2, Y: 2, Orientation: "N", Planet: planet}

	fmt.Println("Welcome to the Mars Rover Simulation!")
	fmt.Println("The Rover is initially at position (2, 2) facing North.")
	fmt.Println("You can enter the following commands:")
	fmt.Println("'M' to move forward, 'B' to move backward, 'L' to turn left, 'R' to turn right")
	fmt.Println("Enter 'Q' to quit the simulation.")
	fmt.Println("Please enter a series of commands (e.g., 'MRMLM'): ")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		commands := scanner.Text()

		if strings.ToUpper(commands) == "Q" {
			fmt.Println("Exiting simulation. Thank you!")
			break
		}

		for _, cmd := range strings.ToUpper(commands) {
			switch cmd {
			case 'M', 'B', 'L', 'R':
				myRover.ExecuteCommand(string(cmd))
				myRover.PrintState()
			default:
				fmt.Println("Invalid command. Please enter 'M', 'B', 'L', 'R' or 'Q' to quit.")
			}
		}
	}
}
