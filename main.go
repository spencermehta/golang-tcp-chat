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

  for {
	conn, err := ln.Accept()
	if err != nil {
	  fmt.Printf("error: %s\n", err)
	}

	readMessage(conn)
  }
}

func readMessage(conn net.Conn) {
    msg, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	fmt.Printf("%s", msg)
	conn.Write([]byte("hi"))
}
