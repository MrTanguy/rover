package planet

import (
	"fmt"
	"testing"
)

func TestNewObstacles(t *testing.T) {
	obstacles := NewObstacles()
	if obstacles == nil {
		t.Fatalf("NewObstacles() returned nil")
	}
	if len(obstacles.obstacles) != 0 {
		t.Errorf("Expected empty obstacles map, got %d elements", len(obstacles.obstacles))
	}
}

func TestSet(t *testing.T) {
	obstacles := NewObstacles()

	// Test setting a new obstacle
	err := obstacles.Set(1, 2)
	if err != nil {
		t.Fatalf("Set(1, 2) returned an error: %v", err)
	}

	// Verify obstacle is set
	if !obstacles.Check(1, 2) {
		t.Errorf("Check(1, 2) returned false, expected true")
	}

	// Test setting an already existing obstacle
	err = obstacles.Set(1, 2)
	if err == nil {
		t.Errorf("Expected error for setting an already existing obstacle, got nil")
	} else if err.Error() != fmt.Sprintf(ErrAlreadySet, 1, 2) {
		t.Errorf("Expected error '%s', got '%v'", fmt.Sprintf(ErrAlreadySet, 1, 2), err)
	}
}

func TestCheck(t *testing.T) {
	obstacles := NewObstacles()

	// Check obstacle that hasn't been set
	if obstacles.Check(3, 4) {
		t.Errorf("Check(3, 4) returned true, expected false")
	}

	// Set an obstacle and verify
	obstacles.Set(3, 4)
	if !obstacles.Check(3, 4) {
		t.Errorf("Check(3, 4) returned false, expected true")
	}
}
