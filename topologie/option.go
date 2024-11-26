package planet

type PlanetOption func(*Planet)

func WithSize(width int, height int) PlanetOption {
	return func(p *Planet) {
		*p = &FinishedPlanet{
			Width:  width,
			Height: height,
		}
	}
}

func WithObstacles(probability uint) PlanetOption {
	return func(p *Planet) {
		if planet, ok := (*p).(*FinishedPlanet); ok {
			planet.seedObstacles(probability)
		}
	}
}
