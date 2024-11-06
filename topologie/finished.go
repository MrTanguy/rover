package planet

type FinishedPlanet struct {
	Width  int
	Height int
}

func (p *FinishedPlanet) WrapX(x int) int {
	if x < 0 {
		return (x + p.Width) % p.Width
	}
	return x % p.Width
}

func (p *FinishedPlanet) WrapY(y int) int {
	if y < 0 {
		return (y + p.Height) % p.Height
	}
	return y % p.Height
}
