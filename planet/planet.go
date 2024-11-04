package planet

type Planet interface {
	WrapX(x int) int
	WrapY(y int) int
}

func New(opts ...PlanetOption) Planet {
	p := &FinishedPlanet{
		Width:  10,
		Height: 10,
	}

	for _, opt := range opts {
		return opt()
	}

	return p
}
