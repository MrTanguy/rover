package socket

import (
	"net"
	"testing"

	"github.com/MrTanguy/rover/missionControl/interpreter"
	"github.com/MrTanguy/rover/rover"
)

// mockRover used to simulate rover movements
type mockRover struct{}

func (m *mockRover) Forward() rover.State   { return rover.State{} }
func (m *mockRover) Backward() rover.State  { return rover.State{} }
func (m *mockRover) TurnLeft() rover.State  { return rover.State{} }
func (m *mockRover) TurnRight() rover.State { return rover.State{} }

func TestSocketPilotHandleConnection(t *testing.T) {
	r := &mockRover{}
	newInterpreter := interpreter.NewInterpreter(r)
	socketPilot := &SocketPilot{Interpreter: newInterpreter}

	// Create a local in-memory pipe
	client, server := net.Pipe()
	defer func(client net.Conn) {
		err := client.Close()
		if err != nil {
			t.Fatalf("Failed to close client pipe: %v", err)
		}
	}(client)
	defer func(server net.Conn) {
		err := server.Close()
		if err != nil {
			t.Fatalf("Failed to close server pipe: %v", err)
		}
	}(server)

	go socketPilot.handleConnection(server)

	tests := []struct {
		input    string
		expected string
	}{
		{"F", "Moved forward\n"},
		{"B", "Moved backward\n"},
		{"L", "Turned left\n"},
		{"R", "Turned right\n"},
		{"X", "Invalid command. Please enter F, B, L, R or Q to quit.\n"},
	}

	for _, tt := range tests {
		t.Run("Command "+tt.input, func(t *testing.T) {
			_, err := client.Write([]byte(tt.input))
			if err != nil {
				t.Fatalf("Failed to write to pipe: %v", err)
			}

			buf := make([]byte, 1024)
			n, err := client.Read(buf)
			if err != nil {
				t.Fatalf("Failed to read from pipe: %v", err)
			}

			got := string(buf[:n])
			if got != tt.expected {
				t.Errorf("got %q, want %q", got, tt.expected)
			}
		})
	}
}
