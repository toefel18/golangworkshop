package main

import (
	"bufio"
	"fmt"
	"goworkshop/docker_example/common"
	"log"
	"net"
)

func connect() net.Conn {

	var conn net.Conn
	var err error

	for conn, err = net.Dial("tcp", "server:4343"); err != nil; {
		log.Fatal(err)
		common.DelaySecond(1)
	}

	return conn
}

func main() {

	// connect to this socket
	conn := connect()

	log.Println("connected")

	for {

		text := common.RandomString(10)

		// send to socket
		log.Printf("sending: %s\n", text)
		fmt.Fprintf(conn, text+"\n")

		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')

		log.Print("Message from server: " + message)

		common.DelaySecond(5)
	}
}
