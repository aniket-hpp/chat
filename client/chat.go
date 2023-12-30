package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"

	"atomicgo.dev/cursor"
	"github.com/gorilla/websocket"
)

func Chat(conn *websocket.Conn, sc *bufio.Reader) {
	fmt.Println("Start Chat: ")

	for {
		smsg, _ := sc.ReadString('\n')
		if len(strings.TrimSpace(string(smsg))) != 0 {
			err := conn.WriteMessage(websocket.TextMessage, []byte(smsg[:len(smsg)-1]))

			if err != nil {
				log.Printf("error: %s\n", err.Error())
				return
			}

			cursor.ClearLinesUp(1)
			fmt.Printf("Sent: %s\n", smsg[:len(smsg)-1])
		}
	}
}
