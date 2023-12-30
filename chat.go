package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func Chat(w http.ResponseWriter, r *http.Request) {
	conn, err := Upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Printf("err: %s\n", err.Error())
		return
	}

	c := Client{ws: conn}
	c.name = getName(conn)
	if c.name == "" {
		return
	}

	Clients[c] = true
	Sndr = &Client{name: "Connected", ws: conn}
	BoardCast <- c.name

	for {
		_, bytes, err := conn.ReadMessage()
		if err != nil {
			log.Printf("err: %s\n", err.Error())
			conn.Close()
			delete(Clients, c)

			Sndr = &Client{name: "Disconnected"}
			BoardCast <- c.name
			return
		}

		Sndr = &c
		BoardCast <- string(bytes)
	}
}

func getName(conn *websocket.Conn) string {
	_, bytes, err := conn.ReadMessage()
	if err != nil {
		conn.Close()
		return ""
	}

	return string(bytes)
}
