package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

// Exmaple read from socket.
func read() {
	socket := os.Getenv("ENVIROMENTAL_VARIABLE")

	conn, err := net.Dial("unix", "/tmp/file/"+socket+"/.socket2.sock")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

// Exmaple write to socket.
func main() {
	socket := os.Getenv("ENVIRONMENTAL_VARIABLE")

	conn, err := net.Dial("unix", "/tmp/file/"+socket+"/.socket.sock")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("message to write"))
	if err != nil {
		log.Fatal(err)
	}

}
