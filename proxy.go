package main

import (
	"io"
	"log"
	"net"
)

func main() {

	// listen on port 80

	listener, err := net.Listen("tcp", ":80")

	if err != nil {

		log.Fatalln(err)
	}

	for {

		conn, err := listener.Accept()

		if err != nil {

			log.Fatalln("Unable to accept connection")
		}

		go handle(conn)
	}

}

func handle(src net.Conn) {

	dst, err := net.Dial("tcp", "https://nmap.org/:80")

	if err != nil {

		log.Fatalln("Unable to connect to the host")

	}

	defer dst.Close()

	go func() {

		// COpy our src's output to the destination

		if _, err := io.Copy(dst, src); err != nil {

			log.Fatalln(err)

		}

	}()

	// Copy from dst's output to src

	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}

}
