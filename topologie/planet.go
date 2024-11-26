package planet

type Planet interface {
	WrapX(x int) int
	WrapY(y int) int
	seedObstacles(probability uint)
	CheckObstacle(x uint, y uint) bool
}

func New(infinite bool, opts ...PlanetOption) (Planet, error) {
	var p Planet
	p = &FinishedPlanet{
		Width:  10,
		Height: 10,
	}
	if infinite {
		p = &InfinitePlanet{}
	}
	for _, opt := range opts {
		opt(&p)
	}
	return p, nil
}
