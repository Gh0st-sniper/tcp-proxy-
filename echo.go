package main

import (
	"io"
	"log"
	"net"
)

func main() {

	// Listen on port 4000 on all interfaces

	listener, err := net.Listen("tcp", ":4000")

	if err != nil {

		log.Fatalln("Unable to bind to port 4000")
	}

	log.Println("Listening on 0.0.0.:4000")

	for {

		// wait for connection .  Create net.Conn on connection established

		conn, err := listener.Accept()

		log.Println("Connection established")
		if err != nil {

			log.Fatalln("Unable to accept connection")

		}

		go echo(conn)

	}

}

func echo(conn net.Conn) {

	defer conn.Close()

	b := make([]byte, 512)

	for {

		size, err := conn.Read(b[0:])

		if err == io.EOF {

			log.Println("Client disconnected")
			break
		}

		if err != nil {

			log.Println("Unexpected Error")

		}

		log.Printf("RECEIVED %d bytes: %s\n", size, string(b))

		// send data

		log.Println("Writing data")

		if _, err := conn.Write(b[0:size]); err != nil {

			log.Fatalln("Unable to write data")
		}
	}

}
