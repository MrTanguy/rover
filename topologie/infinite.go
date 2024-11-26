package planet

type InfinitePlanet struct {
}

func (p *InfinitePlanet) WrapX(x int) int {
	return 0
}

func (p *InfinitePlanet) WrapY(y int) int {
	return 0
}
