package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	roverBuilder "github.com/MrTanguy/rover/builder"
	"github.com/MrTanguy/rover/rover"
)

func main() {

	var myRover rover.Rover = roverBuilder.Build(roverBuilder.WithPlanet(5, 5), roverBuilder.WithPosition(2, 2))

	fmt.Println("Welcome to the Mars Rover Simulation!")
	fmt.Println("The Rover is initially at position (2, 2) facing North.")
	fmt.Println("You can enter the following commands:")
	fmt.Println("'F' to move forward, 'B' to move backward, 'L' to turn left, 'R' to turn right")
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
			case 'F', 'B', 'L', 'R':
				myRover.ExecuteCommand(string(cmd))
			default:
				fmt.Println("Invalid command. Please enter 'F', 'B', 'L', 'R' or 'Q' to quit.")
			}
		}
	}
}
