package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	connectionChannel := make(chan net.Conn)
	messageChannel := make(chan []byte)
	var connections []net.Conn
	port := "2222"
	tcpAddress, err := net.ResolveTCPAddr("tcp", ":"+port)
	checkServerErr(err)
	listener, err := net.ListenTCP("tcp", tcpAddress)
	checkServerErr(err)
	fmt.Printf("Listening on port %s\n", port)
	go handleConnections(listener, connectionChannel, messageChannel)
	go broadcastMessages(messageChannel, &connections)
	for {
		conn := <-connectionChannel
		connections = append(connections, conn)
		fmt.Println("Client connected: ", conn)
		fmt.Println("Total clients: ", len(connections))
	}
}

func handleConnections(listener *net.TCPListener, connectionChannel chan net.Conn, messageChannel chan []byte) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		go handleConnection(conn, connectionChannel, messageChannel)
	}
}

func handleConnection(conn net.Conn, connectionChannel chan net.Conn, messageChannel chan []byte) {
	connectionChannel <- conn
	daytime := time.Now().String()
	conn.Write([]byte("Hey thanks for showing up! The time is " + daytime + "\n"))
	conn.Write([]byte("\nWhat's your message?\n"))
	for {
		var buf [512]byte
		messageLength, err := conn.Read(buf[0:])
		shouldExit := checkClientErr(err, conn)
		if shouldExit {
			return
		}
		messageChannel <- buf[0:messageLength]
	}
}

func broadcastMessages(messageChannel chan []byte, connections *[]net.Conn) {
	for {
		message := <-messageChannel
		for _, c := range *connections {
			c.Write([]byte("        "))
			c.Write(message)
		}
		fmt.Printf(string(message))
	}
}

func checkServerErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func checkClientErr(err error, client net.Conn) bool {
	if err != nil {
		fmt.Println(err.Error())
		client.Close()
		return true
	}
	return false
}
