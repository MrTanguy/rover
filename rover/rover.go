package rover

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
