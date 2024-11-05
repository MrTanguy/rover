package interpreter

import (
	"bytes"
	"fmt"
	"github.com/MrTanguy/rover/rover"
	"io"
	"os"
	"testing"
)

type mockRover struct{}

func (m *mockRover) Forward() rover.State   { return rover.State{} }
func (m *mockRover) Backward() rover.State  { return rover.State{} }
func (m *mockRover) TurnLeft() rover.State  { return rover.State{} }
func (m *mockRover) TurnRight() rover.State { return rover.State{} }

func TestInterpreterCommands(t *testing.T) {
	r := &mockRover{}
	i := NewInterpreter(r)

	tests := []struct {
		cmd      string
		expected string
	}{
		{"F", "Moved forward\n"},
		{"B", "Moved backward\n"},
		{"L", "Turned left\n"},
		{"R", "Turned right\n"},
		{"X", "Invalid command. Please enter F, B, L, R or Q to quit.\n"},
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
