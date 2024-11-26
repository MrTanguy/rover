package planet

import (
	"fmt"

	"golang.org/x/exp/rand"
)

type FinishedPlanet struct {
	Width    int
	Height   int
	obstacle *Obstacles
}

func (p *FinishedPlanet) WrapX(x int) int {
	if x <= 0 {
		return (x + p.Width) % p.Width
	}
	return x % p.Width
}

func (p *FinishedPlanet) WrapY(y int) int {
	if y <= 0 {
		return (y + p.Height) % p.Height
	}
	return y % p.Height
}

// SeedObstacles seeds the planet with obstacles based on the given probability.
// The probability parameter is a percentage (0-100) that determines the density of obstacles.
// The function calculates the number of obstacles to place based on the planet's size and the given probability.
// If the calculated number of obstacles exceeds the planet's size, it is capped to the maximum possible value.
// Obstacles are placed at random positions on the planet, ensuring no duplicate positions.
func (p *FinishedPlanet) seedObstacles(probability uint) {

	if probability < 0 || probability > 100 {
		fmt.Println("Invalid probability value, setting it to 0")
		probability = 0
	}

	planetSize := p.Width * p.Height

	number := planetSize * int(probability) / 100

	if number > planetSize-1 {
		number = planetSize - 1
		fmt.Printf("Number of obstacles is greater than the planet size setting it to %d\n", planetSize-1)
	}

	p.obstacle = NewObstacles()

	for i := int(0); i < number; i++ {
		x, y := p.randPosition()
		err := p.obstacle.Set(int(x), int(y))
		for err != nil {
			x, y = p.randPosition()
			err = p.obstacle.Set(int(x), int(y))
		}
	}
}

func (p *FinishedPlanet) randPosition() (uint, uint) {
	x := uint(rand.Intn(p.Width))
	y := uint(rand.Intn(p.Height))

	for x == 0 && y == 0 {
		x = uint(rand.Intn(p.Width))
		y = uint(rand.Intn(p.Height))
	}

	return x, y
}

func (p *FinishedPlanet) CheckObstacle(x int, y int) bool {
	return p.obstacle.Check(x, y)
}
