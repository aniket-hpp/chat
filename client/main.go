package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

type Msg struct {
	Name string
	Msg  string
}

var (
	Host = "ws://localhost:6969/chat"
)

func main() {
	conn, _, err := websocket.DefaultDialer.Dial(Host, nil)
	sc := bufio.NewReader(os.Stdin)
	if err != nil {
		log.Fatal("error: ", err)
	}

	defer conn.Close()
	err = setName(conn, sc)
	if err != nil {
		log.Fatal("error: ", err)
	}

	go ReadMsg(conn)
	Chat(conn, sc)
}

func setName(conn *websocket.Conn, sc *bufio.Reader) error {
	fmt.Print("Enter your name: ")
	name, _ := sc.ReadString('\n')

	err := conn.WriteMessage(websocket.TextMessage, []byte(name[:len(name)-1]))
	return err
}
