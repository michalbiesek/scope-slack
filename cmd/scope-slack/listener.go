package main

import (
	"fmt"
	"net"
	"os"
)

func listener(scopeServer string) {
	// log.Debug().Msgf("Scope-slack will listen data")
	l, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on localhost:9999")
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	// Send a response back to person contacting us.
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()
}
