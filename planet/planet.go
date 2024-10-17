package planet

type Planet struct {
	Width  int
	Height int
}

func NewPlanet(width, height int) *Planet {
	return &Planet{Width: width, Height: height}
}

func (p *Planet) WrapX(x int) int {
	if x < 0 {
		return (x + p.Width) % p.Width
	}
	return x % p.Width
}

func (p *Planet) WrapY(y int) int {
	if y < 0 {
		return (y + p.Height) % p.Height
	}
	return y % p.Height
}
