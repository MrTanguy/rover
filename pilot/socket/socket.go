package socket

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/MrTanguy/rover/pilot"
	"github.com/MrTanguy/rover/pilot/interpreter"
	"github.com/MrTanguy/rover/rover"
)

// SocketPilot structure to implement Pilot interface and handle socket communication
type SocketPilot struct {
	interpreter.Interpreter
}

// NewSocketPilot creates a new SocketPilot instance
func NewSocketPilot(r rover.Rover) pilot.Pilot {
	return &SocketPilot{
		Interpreter: interpreter.NewInterpreter(r),
	}
}

// Run method to start the socket server and handle incoming commands
func (s *SocketPilot) Run() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer ln.Close()
	fmt.Println("Socket server started on port 8080")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
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

// interpretAndCaptureOutput captures output of Interpreter.Interprete and returns it as a string
func (s *SocketPilot) interpretAndCaptureOutput(commands string) string {
	// Create a pipe to capture output
	r, w, _ := os.Pipe()

	// Save the original os.Stdout
	originalStdout := os.Stdout
	defer func() { os.Stdout = originalStdout }() // Restore os.Stdout after the function

	// Set os.Stdout to the write end of the pipe
	os.Stdout = w

	var outputBuffer bytes.Buffer
	done := make(chan struct{})

	// Start a goroutine to read from the read end of the pipe
	go func() {
		_, err := io.Copy(&outputBuffer, r)
		if err != nil {
			fmt.Println("Error copying from pipe:", err)
		}
		close(done)
	}()

	// Interpret each command
	for _, cmd := range commands {
		s.Interpreter.Interprete(string(cmd))
	}

	// Close the write end of the pipe and wait for the goroutine to finish
	err := w.Close()
	if err != nil {
		fmt.Println("Error closing pipe:", err)
	}
	<-done

	// Return the captured output as a string
	return outputBuffer.String()
}
