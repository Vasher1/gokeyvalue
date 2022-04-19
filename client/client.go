package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please choose udp, tcp, or load")
		return
	}

	if arguments[1] == "tcp" {
		sendCustomTCPRequest()
	} else if arguments[1] == "udp" {
		sendCustomUDPRequest()
	} else if arguments[1] == "load" {
		loadTest()
	}
}

func sendCustomUDPRequest() {
	c, err := net.Dial("udp", "localhost:8081")
	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter request string \n")
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	_, err = fmt.Fprintf(c, input)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("\n Request sent \n")
}

func sendCustomTCPRequest() {
	c, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter request string \n")
	input, _ := reader.ReadString('\n')
	_, err = fmt.Fprintf(c, input)

	if err != nil {
		fmt.Println(err)
	}

	response, _ := bufio.NewReader(c).ReadString('\n')
	fmt.Print("\n Response: \n" + response + "\n")
}

func loadTest() {
	var wg sync.WaitGroup
	var requestsToDo = 1000

	wg.Add(requestsToDo)
	fmt.Printf("Starting \n")

	for i := 1; i < requestsToDo+1; i++ {
		go sendGenericTCPRequest(&wg, i)
	}

	//for i := 1; i < requestsToDo+1; i++ {
	//	go sendHttpRequest(&wg, i)
	//}

	wg.Wait()
	fmt.Printf("Done \n")
}

func sendGenericTCPRequest(wg *sync.WaitGroup, i int) {
	CONNECT := "localhost:8081"

	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Printf(fmt.Sprintf("failed request number %d \n", i))
		wg.Done()
		return
	}

	fmt.Fprintf(c, "HELLO \n")
	reply, _ := bufio.NewReader(c).ReadString('\n')
	fmt.Print("Request #" + strconv.Itoa(i) + " : " + reply + "\n")

	wg.Done()
}

func sendGenericHttpRequest(wg *sync.WaitGroup, i int) {
	CONNECT := "localhost:8081"

	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Printf(fmt.Sprintf("failed request number %d \n", i))
		wg.Done()
		return
	}

	fmt.Fprintf(c, "HELLO \n")
	//reply, _ := bufio.NewReader(c).ReadString('\n')
	//fmt.Print("->: " + reply + "\n")

	fmt.Printf(fmt.Sprintf("completed request number %d \n", i))
	wg.Done()
}
