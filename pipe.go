package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

// /tmp/hypr/[HIS]/.socket2.socket
// Events socket

// /tmp/hypr/[HIS]/.socket.socket
// Request socket

// Exmaple read from socket
func read() {
	socket := os.Getenv("HYPRLAND_INSTANCE_SIGNATURE")

	conn, err := net.Dial("unix", "/tmp/hypr/"+socket+"/.socket2.sock")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

// Exmaple write to socket
func main() {
	socket := os.Getenv("HYPRLAND_INSTANCE_SIGNATURE")

	conn, err := net.Dial("unix", "/tmp/hypr/"+socket+"/.socket.sock")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("dispatch exec alacritty"))
	if err != nil {
		log.Fatal(err)
	}

}
