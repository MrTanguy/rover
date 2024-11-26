package planet_test

import (
	"testing"

	planet "github.com/MrTanguy/rover/topologie"
	"github.com/stretchr/testify/assert"
)

func TestWrapX(t *testing.T) {
	planet, _ := planet.New(false, planet.WithSize(5, 5))

	assert.Equal(t, 4, planet.WrapX(4))  // Position normale
	assert.Equal(t, 0, planet.WrapX(5))  // Doit se replier sur 0
	assert.Equal(t, 4, planet.WrapX(-1)) // -1 doit se replier sur 4
}

func TestWrapY(t *testing.T) {
	planet, _ := planet.New(false, planet.WithSize(5, 5))

	assert.Equal(t, 4, planet.WrapY(4))  // Position normale
	assert.Equal(t, 0, planet.WrapY(5))  // Doit se replier sur 0
	assert.Equal(t, 4, planet.WrapY(-1)) // -1 doit se replier sur 4
}

func TestNewPlanet(t *testing.T) {
	_, err := planet.New(false, planet.WithSize(5, 5))

	assert.Nil(t, err)

	_, err = planet.New(true)
	assert.Nil(t, err)

}
