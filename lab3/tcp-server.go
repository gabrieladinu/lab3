package main

import "net"
import "fmt"
import (
	"bufio"
	//"strings"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Launching server...")

	// listen on all interfaces

	ln, _ := net.Listen("tcp", ":8081")
	ln2, _ := net.Listen("tcp", ":8080")


	// accept connection on port
	conn, _ := ln.Accept()
	conn2, _ := ln2.Accept()

	aux:=0

	// run loop forever (or until ctrl-c)
	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')

		b := strings.TrimSuffix(string(message), "\n")
		i:=0
		i, _ = strconv.Atoi(b)
		fmt.Println(b)

		nr := aux + i
		fmt.Println(nr)

		message2, _ := bufio.NewReader(conn2).ReadString('\n')


		b1 := strings.TrimSuffix(string(message2), "\n")
		j:=0
		j, _ = strconv.Atoi(b1)



		fmt.Println(b1)
		nr = nr +aux + j
		fmt.Println(nr)

		// output message received
		//fmt.Print("Message Received:", string(message))
		// sample process for string received

		//b := strings.TrimSuffix(string(message), "\n")
		//
		//r := []rune(b)
		//for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		//	r[i], r[j] = r[j], r[i]
		//}

		//ceva := strconv.Itoa(nr)

fmt.Println("suma este : "  )
fmt.Println(nr)
		ceva := strconv.Itoa(nr)

fmt.Println("ceva este :")
fmt.Println(ceva)
		conn.Write([]byte(string(ceva)+ "\n"))

		conn2.Write([]byte(string(ceva)+ "\n"))

	}
}
