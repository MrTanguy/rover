package main

import (
	"fmt"

	"github.com/MrTanguy/rover/missionControl/socket"

	roverBuilder "github.com/MrTanguy/rover/builder"
	"github.com/MrTanguy/rover/rover"
)

const (
	positionX = 2
	positionY = 2
)

func main() {

	var myRover rover.Rover = roverBuilder.Build(roverBuilder.WithPlanet(5, 5), roverBuilder.WithPosition(positionX, positionY))

	fmt.Println("Welcome to the Mars Rover Simulation!")
	fmt.Printf("The Rover is initially at position (%d, %d) facing North. \n", positionX, positionY)
	fmt.Println("You can enter the following commands:")
	fmt.Println("'F' to move forward, 'B' to move backward, 'L' to turn left, 'R' to turn right")
	fmt.Println("Enter 'Q' to quit the simulation.")
	fmt.Println("Please enter a series of commands (e.g., 'MRMLM'): ")

	interpreter := socket.NewSocketPilot(myRover)
	interpreter.Run()

}
