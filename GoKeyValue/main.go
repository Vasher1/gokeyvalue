package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {
	controller := newKeyValueHttpController()
	http.Handle("/keyvalue", controller)
	http.Handle("/keyvalue/", controller)
	http.ListenAndServe(":8080", nil)

	// Listen for incoming connections.
	addr := "localhost:8888"
	l, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer l.Close()
	host, port, err := net.SplitHostPort(l.Addr().String())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Listening on host: %s, port: %s\n", host, port)

	for {
		// Listen for an incoming connection
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		// Handle connections in a new goroutine
		go func(conn net.Conn) {
			buf := make([]byte, 1024)
			len, err := conn.Read(buf)
			if err != nil {
				fmt.Printf("Error reading: %#v\n", err)
				return
			}
			fmt.Printf("Message received: %s\n", string(buf[:len]))

			conn.Write([]byte("Message received.\n"))
			conn.Close()
		}(conn)
	}
}
