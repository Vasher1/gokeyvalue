package tcpServer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"keyvaluepairexercise/keyValueService"
	"net"
	"os"
)

func TcpListen() {
	l, err := net.Listen("tcp", "localhost:8081")

	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	// Read message
	buf := make([]byte, 4096)
	_, err := conn.Read(buf)

	if err != nil {
		conn.Write([]byte("Error reading request:" + err.Error()))
		conn.Close()
		return
	}

	// Parse request
	var r request
	var response string
	err = json.Unmarshal(bytes.Trim(buf, "\x00"), &r)
	if err != nil {
		conn.Write([]byte("Error unmarshalling json:" + err.Error()))
		conn.Close()
		return
	}

	switch r.Method {
	case "GETALL":
		response = getAll(r)
	case "GET":
		response = get(r)
	case "PUT":
		response = put(r)
	case "DELETE":
		response = delete(r)
	default:
		response = "ERROR"
	}

	conn.Write([]byte(response))
	conn.Close()
}

// Methods

func getAll(r request) string {
	err, store := keyValueService.ReadAll()

	if err != nil {
		return err.Error()
	}

	jsonValue, err := json.Marshal(store)

	if err != nil {
		return err.Error()
	}

	return string(jsonValue)
}

func get(r request) string {
	err, value := keyValueService.Read(r.Key)

	if err != nil {
		return err.Error()
	}

	jsonValue, err := json.Marshal(value)

	if err != nil {
		return err.Error()
	}

	return string(jsonValue)
}

func put(r request) string {
	err := keyValueService.Add(r.Key, r.Value)

	if err != nil {
		return err.Error()
	}

	return "New key value pair added succesfully"
}

func delete(r request) string {
	err := keyValueService.Remove(r.Key)

	if err != nil {
		return err.Error()
	}

	return "Key value pair deleted succesfully"
}

// Structs

type request struct {
	Key    string
	Value  interface{}
	Method string
}
