package planet

import (
	"testing"
)

func TestFinishedWrapX(t *testing.T) {
	planet := &FinishedPlanet{Width: 10}
	tests := []struct {
		x      int
		expect int
	}{
		{-1, 9},
		{0, 0},
		{10, 0},
		{12, 2},
	}

	for _, test := range tests {
		result := planet.WrapX(test.x)
		if result != test.expect {
			t.Errorf("WrapX(%d) = %d; want %d", test.x, result, test.expect)
		}
	}
}

func TestFinishedWrapY(t *testing.T) {
	planet := &FinishedPlanet{Height: 10}
	tests := []struct {
		y      int
		expect int
	}{
		{-1, 9},
		{0, 0},
		{10, 0},
		{15, 5},
	}

	for _, test := range tests {
		result := planet.WrapY(test.y)
		if result != test.expect {
			t.Errorf("WrapY(%d) = %d; want %d", test.y, result, test.expect)
		}
	}
}

func TestSeedObstacles(t *testing.T) {
	planet := &FinishedPlanet{Width: 10, Height: 10}

	// Seed with 50% probability
	planet.seedObstacles(50)

	// Total obstacles should be close to 50% of planet size
	planetSize := planet.Width * planet.Height
	expectedObstacles := planetSize * 50 / 100

	count := 0
	for x := 0; x < planet.Width; x++ {
		for y := 0; y < planet.Height; y++ {
			if planet.CheckObstacle(x, y) {
				count++
			}
		}
	}

	if count != expectedObstacles {
		t.Errorf("Expected approximately %d obstacles, got %d", expectedObstacles, count)
	}
}

func TestFinishedSeedObstaclesInvalidProbability(t *testing.T) {
	planet := &FinishedPlanet{Width: 10, Height: 10}

	// Seed with invalid probability
	planet.seedObstacles(150)

	// Ensure no obstacles were placed
	count := 0
	for x := 0; x < planet.Width; x++ {
		for y := 0; y < planet.Height; y++ {
			if planet.CheckObstacle(x, y) {
				count++
			}
		}
	}

	if count != 0 {
		t.Errorf("Expected no obstacles, but found %d", count)
	}
}

func TestCheckObstacle(t *testing.T) {
	planet := &FinishedPlanet{Width: 10, Height: 10}
	planet.seedObstacles(100)

	// Ensure CheckObstacle returns true for obstacles
	x, y := 1, 1
	planet.obstacle.Set(x, y)

	if !planet.CheckObstacle(x, y) {
		t.Errorf("CheckObstacle(%d, %d) returned false, expected true", x, y)
	}

	// Ensure CheckObstacle returns false for non-obstacles
	if planet.CheckObstacle(0, 0) {
		t.Errorf("CheckObstacle(0, 0) returned true, expected false")
	}
}

func TestRandPosition(t *testing.T) {
	planet := &FinishedPlanet{Width: 10, Height: 10}

	for i := 0; i < 100; i++ {
		x, y := planet.randPosition()
		if x < 0 || int(x) >= planet.Width || y < 0 || int(y) >= planet.Height {
			t.Errorf("randPosition() returned invalid position (%d, %d)", x, y)
		}
	}
}
