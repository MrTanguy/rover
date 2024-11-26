package socket

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"strings"

	"github.com/MrTanguy/rover/communication/interpreter"
	"github.com/MrTanguy/rover/missionControl"
	"github.com/MrTanguy/rover/rover"
)

// SocketPilot structure to implement Pilot interface and handle socket communication
type SocketPilot struct {
	interpreter.Interpreter
}

// NewSocketPilot creates a new SocketPilot instance
func NewSocketPilot(r rover.Rover) missionControl.MissionControl {
	return &SocketPilot{
		Interpreter: interpreter.NewInterpreter(r),
	}
}

// Run method to start the socket server and handle incoming commands
func (s *SocketPilot) Run() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Error starting server: %w", err)
		return
	}
	defer ln.Close()
	fmt.Println("Socket server started on port 8080")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %w", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

// handleConnection reads commands from the socket and passes them to the interpreter
func (s *SocketPilot) handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}

		commands := string(buffer[:n])
		response := s.interpretAndCaptureOutput(commands)
		_, err = conn.Write([]byte(response)) // Send response back to the client
		if err != nil {
			fmt.Println("Error writing to connection:", err)
			return
		}
	}
}

func (s *SocketPilot) interpretAndCaptureOutput(commands string) string {
	r, w, _ := os.Pipe()

	originalStdout := os.Stdout
	defer func() { os.Stdout = originalStdout }()

	os.Stdout = w

	var outputBuffer bytes.Buffer
	done := make(chan struct{})

	go func() {
		_, err := io.Copy(&outputBuffer, r)
		if err != nil {
			fmt.Println("Error copying from pipe:", err)
		}
		close(done)
	}()

	for _, cmd := range commands {
		if strings.TrimSpace(string(cmd)) == "" {
			continue
		}
		s.Interpreter.Interprete(strings.ToUpper(string(cmd)))
	}

	err := w.Close()
	if err != nil {
		fmt.Println("Error closing pipe:", err)
	}
	<-done

	return outputBuffer.String()
}
