package main

import (
	"bufio"
	"fmt"
	"net"
)

type client struct {
	conn net.Conn
	messages chan<- message
}

func (client *client) readMessages() {
	for {
    	msg, err := bufio.NewReader(client.conn).ReadString('\n')
		if err != nil {
			fmt.Printf("error: %s", err)
		}
		fmt.Printf("%s: %s", client.conn.RemoteAddr(), msg)

		client.messages <- message {
			client: client,
			msg: msg,
		}
	}
}
