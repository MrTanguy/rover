package planet

import "fmt"

type Obstacles struct {
	obstacles map[int]map[int]bool
}

const (
	ErrAlreadySet = "obstacle already set (%d, %d)"
)

func NewObstacles() *Obstacles {
	return &Obstacles{
		obstacles: make(map[int]map[int]bool),
	}
}

func (p *Obstacles) Set(x int, y int) error {
	// Si la clé x n'existe pas encore dans la map, initialisez-la
	if _, ok := p.obstacles[x]; !ok {
		p.obstacles[x] = make(map[int]bool)
	}
	// Si la position (x, y) est déjà définie, renvoyez une erreur
	if p.obstacles[x][y] {
		return fmt.Errorf(ErrAlreadySet, x, y)
	}
	// Sinon, définissez l'obstacle
	p.obstacles[x][y] = true
	return nil
}

func (p *Obstacles) Check(x int, y int) bool {
	ok, _ := p.obstacles[x][y]
	return ok
}
