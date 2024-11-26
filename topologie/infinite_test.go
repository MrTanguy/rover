package planet

import (
	"testing"
)

// Test WrapX
func TestInfiniteWrapX(t *testing.T) {
	planet := &InfinitePlanet{}
	x := 42
	if wrappedX := planet.WrapX(x); wrappedX != x {
		t.Errorf("WrapX(%d) = %d, want %d", x, wrappedX, x)
	}
}

// Test WrapY
func TestInfiniteWrapY(t *testing.T) {
	planet := &InfinitePlanet{}
	y := 99
	if wrappedY := planet.WrapY(y); wrappedY != y {
		t.Errorf("WrapY(%d) = %d, want %d", y, wrappedY, y)
	}
}

// Test seedObstacles with valid probability
func TestSeedObstaclesValidProbability(t *testing.T) {
	planet := &InfinitePlanet{}
	planet.seedObstacles(50)

	if planet.probability != 50 {
		t.Errorf("Expected probability to be 50, got %d", planet.probability)
	}

	if planet.obstacles == nil {
		t.Errorf("Expected obstacles to be initialized, got nil")
	}
}

// Test seedObstacles with invalid probability
func TestInfiniteSeedObstaclesInvalidProbability(t *testing.T) {
	planet := &InfinitePlanet{}
	planet.seedObstacles(120) // Invalid probability > 100

	if planet.probability != 0 {
		t.Errorf("Expected probability to be set to 0 for invalid input, got %d", planet.probability)
	}
}

// Test CheckObstacle without obstacles
func TestCheckObstacleNoObstacles(t *testing.T) {
	planet := &InfinitePlanet{}
	planet.seedObstacles(0) // Set probability to 0 to avoid random obstacles

	x, y := 1, 1
	if planet.CheckObstacle(x, y) {
		t.Errorf("CheckObstacle(%d, %d) returned true, expected false", x, y)
	}
}

// Test CheckObstacle with random obstacle generation
func TestCheckObstacleWithProbability(t *testing.T) {
	planet := &InfinitePlanet{}
	planet.seedObstacles(100) // Set probability to 100 to always generate obstacles

	x, y := 2, 2
	if !planet.CheckObstacle(x, y) {
		t.Errorf("CheckObstacle(%d, %d) returned false, expected true", x, y)
	}
}

// Test CheckObstacle re-checking same obstacle
func TestCheckObstacleReCheck(t *testing.T) {
	planet := &InfinitePlanet{}
	planet.seedObstacles(50)

	x, y := 3, 3
	planet.obstacles.Set(x, y) // Manually set obstacle

	if !planet.CheckObstacle(x, y) {
		t.Errorf("CheckObstacle(%d, %d) returned false, expected true", x, y)
	}
}
