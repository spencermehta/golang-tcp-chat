package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
  ln, err := net.Listen("tcp", ":8080")
  if err != nil {
	fmt.Printf("error: %s\n", err)
  }

  defer ln.Close()
  for {
	conn, err := ln.Accept()
	if err != nil {
	  fmt.Printf("error: %s\n", err)
	}

	go readMessage(conn)
  }
}

func readMessage(conn net.Conn) {
	for {
    	msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("error: %s", err)
		}
		fmt.Printf("%s: %s", conn.RemoteAddr(), msg)
		conn.Write([]byte("hi\n"))
	}
}
