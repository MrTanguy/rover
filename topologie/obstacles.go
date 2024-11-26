package planet

import "fmt"

type Obstacles struct {
	obstacles map[uint]map[uint]bool
}

const (
	ErrAlreadySet = "obstacle already set (%d, %d)"
)

func NewObstacles() *Obstacles {
	return &Obstacles{
		obstacles: make(map[uint]map[uint]bool),
	}
}

func (p *Obstacles) Set(x uint, y uint) error {
	if _, ok := p.obstacles[x]; ok {
		return fmt.Errorf(ErrAlreadySet, x, y)
	}
	p.obstacles[x][y] = true
	return nil
}

func (p *Obstacles) Check(x uint, y uint) bool {
	ok, _ := p.obstacles[x][y]
	return ok
}
