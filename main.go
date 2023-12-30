package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Client struct {
	ws   *websocket.Conn
	name string
}

const (
	PORT = "6969"
)

var (
	Upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	Clients   = make(map[Client]bool)
	BoardCast = make(chan string)
	Sndr      *Client
)

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/chat", Chat)

	fmt.Println("LOG: Server Started at PORT:", PORT)
	go BoardCastMsg()

	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		log.Fatal(err)
	}
}
