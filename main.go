package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var (
	listenFlag = flag.String("listen", ":5678", "address and port to listen")
	textFlag   = flag.String("text", "", "text to respond")

	// stdoutW and stderrW are for overriding in test.
	stdoutW = os.Stdout
	stderrW = os.Stderr
)

const (
	tcpLogDateFormat string = "2006/01/02 15:04:05"
	tcpLogFormat     string = "%v [remote] %s %s [local] %s %s [duration] %v\n"
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

	// Create listener
	ln, err := net.Listen("tcp", *listenFlag)
	if err != nil {
		log.Fatalf("[ERR] failed to create listener: %s", err)
	}
	log.Printf("[INFO] server is listening on %s\n", *listenFlag)

	// Serve requests
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("[ERR] failed to accept connection: %s", err)
		}
		go handleConnection(conn, *textFlag)
	}
}

func handleConnection(conn net.Conn, v string) {
	// Log connection info
	defer func(start time.Time) {
		end := time.Now()
		dur := end.Sub(start)
		fmt.Fprintf(stdoutW, tcpLogFormat,
			end.Format(tcpLogDateFormat),
			conn.RemoteAddr().String(), conn.RemoteAddr().Network(),
			conn.LocalAddr().String(), conn.LocalAddr().Network(),
			dur)
	}(time.Now())

	// Respond with text
	fmt.Fprintln(conn, v)
	conn.Close()
}
