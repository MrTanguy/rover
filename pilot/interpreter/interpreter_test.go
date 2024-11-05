package interpreter

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/MrTanguy/rover/rover"
)

type mockRover struct{}

func (m *mockRover) Forward() rover.State {
	return rover.State{
		X:           0,
		Y:           0,
		Orientation: "N",
	}
}
func (m *mockRover) Backward() rover.State {
	return rover.State{
		X:           1,
		Y:           1,
		Orientation: "S",
	}
}
func (m *mockRover) TurnLeft() rover.State {
	return rover.State{
		X:           2,
		Y:           2,
		Orientation: "W",
	}
}
func (m *mockRover) TurnRight() rover.State {
	return rover.State{
		X:           3,
		Y:           3,
		Orientation: "E",
	}
}

func TestInterpreterCommands(t *testing.T) {
	r := &mockRover{}
	i := NewInterpreter(r)

	tests := []struct {
		cmd      string
		expected string
	}{
		{"F", "rover is at x: 0 y: 0 orientation: N\n"},
		{"B", "rover is at x: 1 y: 1 orientation: S\n"},
		{"L", "rover is at x: 2 y: 2 orientation: W\n"},
		{"R", "rover is at x: 3 y: 3 orientation: E\n"},
		{"X", "Invalid command. Please enter F, B, L, R or Q to quit. \n"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Command %s", tt.cmd), func(t *testing.T) {
			// Create a pipe to capture output
			r, w, _ := os.Pipe()

			// Save the original os.Stdout and defer its restoration
			originalStdout := os.Stdout
			defer func() { os.Stdout = originalStdout }()
			os.Stdout = w

			// Capture the output in a buffer
			var buf bytes.Buffer
			done := make(chan struct{})
			go func() {
				_, err := io.Copy(&buf, r)
				if err != nil {
					t.Errorf("Error copying from pipe: %v", err)
					return
				}
				close(done)
			}()

			// Run the command
			i.Interprete(tt.cmd)

			// Close the writer and wait for the copying to complete
			err := w.Close()
			if err != nil {
				t.Fatalf("Error closing pipe: %v", err)
			}
			<-done

			// Check if the output matches the expected result
			got := buf.String()
			if got != tt.expected {
				t.Errorf("got %q, want %q", got, tt.expected)
			}
		})
	}
}
