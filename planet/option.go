package planet

type PlanetOption func() Planet

func WithInfity() PlanetOption {
	return func() Planet {
		return &InfinitePlanet{}
	}
}

func WithSize(width int, height int) PlanetOption {
	return func() Planet {
		return &FinishedPlanet{
			Width:  width,
			Height: height,
		}
	}
}