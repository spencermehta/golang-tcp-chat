package main

import (
	"fmt"
	"net"
)

type server struct {
	clients map[net.Addr]net.Conn
	messages chan message
}

func (server *server) run() {
	for msg := range server.messages {
		server.sendMessage(msg)
	}
}

func createServer() *server {
	return &server {
		clients: make(map[net.Addr]net.Conn),
		messages: make(chan message),
	}
}

func (server *server) addClient(conn net.Conn) {
	server.clients[conn.RemoteAddr()] = conn

	client := &client {
		conn: conn,
		messages: server.messages,
	}

	go client.readMessages()
}

func (server *server) sendMessage(msg message) {
	msgToSend := fmt.Sprintf("[%s] %s", msg.client.conn.RemoteAddr(), msg.msg)
	for _, client := range server.clients {
		if client == msg.client.conn {
			continue
		}
		client.Write([]byte(msgToSend))
	}
}
