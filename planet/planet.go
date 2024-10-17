package planet

// Planet represents a toroidal planet with defined width and height
type Planet struct {
	Width  int
	Height int
}

// NewPlanet creates a new Planet with the given width and height
func NewPlanet(width, height int) *Planet {
	return &Planet{Width: width, Height: height}
}

// WrapX wraps the X coordinate around the planet horizontally (toroidal)
func (p *Planet) WrapX(x int) int {
	if x < 0 {
		return (x + p.Width) % p.Width
	}
	return x % p.Width
}

// WrapY wraps the Y coordinate around the planet vertically (toroidal)
func (p *Planet) WrapY(y int) int {
	if y < 0 {
		return (y + p.Height) % p.Height
	}
	return y % p.Height
}
