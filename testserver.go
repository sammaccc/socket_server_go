package main

import (
	"bufio"
	"fmt"
	"net"
)

var maxPlayers int = 5
var serverPassword string = "gummybears\n"
var numPlayers int = 0

type player struct {
	conn      net.Conn
	playerNum int
}

func read_msg(conn net.Conn, players []player) {
	buf := bufio.NewReader(conn)
	msg, err := buf.ReadString('\n')

	if err != nil {
		fmt.Printf("Client disconnected.\n")

	} else {
		if msg == serverPassword {
			newPlayer := player{conn, numPlayers}
			players = append(players, newPlayer)
			numPlayers += 1
		} else {
			fmt.Printf("%s\n", msg)
			conn.Write([]byte("wrong password bye\n"))
			conn.Close()
		}
	}
}

func client_io(clientConns chan net.Conn, players []player) {

	for {
		conn := <-clientConns

		if numPlayers < maxPlayers {
			conn.Write([]byte("pls enter server password for entry\n"))
			go read_msg(conn, players)
		} else {
			conn.Write([]byte("sorry server full\n"))
			conn.Close()
		}
	}
}

func handle_connection(ln net.Listener, clientConns chan net.Conn) {

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("error accepting\n")
		} else {
			fmt.Printf("accepted connection\n")
			clientConns <- conn
		}
	}

}

func main() {

	clientConns := make(chan net.Conn, 100)
	ln, _ := net.Listen("tcp", ":8080")

	var players []player = make([]player, maxPlayers)

	go handle_connection(ln, clientConns)
	go client_io(clientConns, players)
	for {
		// do something
	}
}
