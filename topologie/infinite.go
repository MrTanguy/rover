package planet

import (
	"fmt"
	"math/rand"
)

type InfinitePlanet struct {
	probability uint
	obstacles   *Obstacles
}

func (p *InfinitePlanet) WrapX(x int) int {
	return x
}

func (p *InfinitePlanet) WrapY(y int) int {
	return y
}

func (p *InfinitePlanet) seedObstacles(probability uint) {
	p.obstacles = NewObstacles()

	if probability > 100 || probability < 0 {
		fmt.Println("Invalid probability value, setting it to 0")
		probability = 0
	}

	p.probability = probability
}

func (p *InfinitePlanet) CheckObstacle(x int, y int) bool {
	a := rand.Intn(100)

	if p.obstacles.Check(x, y) {
		return true
	}

	if a < int(p.probability) {
		p.obstacles.Set(x, y)
		return true
	}
	return false
}
