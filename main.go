package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	// Listen
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer li.Close()

	for {
		// Accept
		conn, err := li.Accept()
		if err != nil {
			panic(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		// Convert text to lower
		ln := strings.ToLower(scanner.Text())
		bs := []byte(ln)
		// Rotate function with the new string
		r := rot13(bs)

		fmt.Fprintf(conn, "%s - %s \n \n", ln, r)
	}
}

// func rot13 => Rotate 13 positions
func rot13(bs []byte) []byte {
	var r13 = make([]byte, len(bs))
	for i, v := range bs {
		if v <= 109 {
			r13[i] = v + 13
		} else {
			r13[i] = v - 13
		}
	}
	return r13
}
