package main

import (
	"fmt"
	"net"
)

func main() {
  s := createServer()

  ln, err := net.Listen("tcp", ":8080")
  if err != nil {
	fmt.Printf("error: %s\n", err)
  }

  go s.run()

  defer ln.Close()
  for {
	conn, err := ln.Accept()
	if err != nil {
	  fmt.Printf("error: %s\n", err)
	}

	s.addClient(conn)
  }
}
