package planet

type InfinitePlanet struct {
}

func (p *InfinitePlanet) WrapX(x int) int {
	return x
}

func (p *InfinitePlanet) WrapY(y int) int {
	return y
}
