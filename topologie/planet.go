package planet

import (
	"errors"
)

type Planet interface {
	WrapX(x int) int
	WrapY(y int) int
}

func New(opts ...PlanetOption) (Planet, error) {
	for _, opt := range opts {
		return opt(), nil
	}
	return nil, errors.New("planet: no valid options provided")
}
