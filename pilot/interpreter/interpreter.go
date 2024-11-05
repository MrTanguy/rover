package interpreter

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MrTanguy/rover/pilot"
	"github.com/MrTanguy/rover/rover"
)

const (
	cmdForward  = "F"
	cmdBackward = "B"
	cmdLeft     = "L"
	cmdRight    = "R"
	cmdQuit     = "Q"
)

type Interpreter struct {
	rover rover.Rover
}

func NewPilot(r rover.Rover) pilot.Pilot {
	return &Interpreter{rover: r}
}

func NewInterpreter(r rover.Rover) Interpreter {
	return Interpreter{rover: r}
}

func (i *Interpreter) Run() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		commands := scanner.Text()
		for _, cmd := range strings.ToUpper(commands) {
			i.Interprete(string(cmd))
		}
	}
}

func (i *Interpreter) Interprete(cmd string) {
	switch cmd {
	case cmdForward:
		fmt.Println(i.rover.Forward())
	case cmdBackward:
		fmt.Println(i.rover.Backward())
	case cmdLeft:
		fmt.Println(i.rover.TurnLeft())
	case cmdRight:
		fmt.Println(i.rover.TurnRight())
	case cmdQuit:
		fmt.Println("Exiting simulation. Thank you!")
		os.Exit(0)
	default:
		fmt.Printf("Invalid command. Please enter %s, %s, %s, %s or %s to quit. \n", cmdForward, cmdBackward, cmdLeft, cmdRight, cmdQuit)
	}
}
