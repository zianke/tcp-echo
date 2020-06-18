package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

var (
	listenFlag = flag.String("listen", ":5678", "address and port to listen")
	textFlag   = flag.String("text", "", "text to respond")

	// stdoutW and stderrW are for overriding in test.
	stdoutW = os.Stdout
	stderrW = os.Stderr
)

func main() {
	flag.Parse()

	// Validation
	if *textFlag == "" {
		fmt.Fprintln(stderrW, "Missing -text option!")
		os.Exit(127)
	}

	args := flag.Args()
	if len(args) > 0 {
		fmt.Fprintln(stderrW, "Too many arguments!")
		os.Exit(127)
	}

	ln, err := net.Listen("tcp", *listenFlag)
	if err != nil {
		log.Fatalf("[ERR] failed to create listener: %s", err)
	}
	log.Printf("[INFO] server is listening on %s\n", *listenFlag)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("[ERR] failed to accept connection: %s", err)
		}
		go handleConnection(conn, *textFlag)
	}
}

func handleConnection(conn net.Conn, v string) {
	fmt.Fprintln(conn, v)
	conn.Close()
}
