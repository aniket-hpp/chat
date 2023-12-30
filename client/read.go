package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func ReadMsg(conn *websocket.Conn) {
	defer conn.Close()
	var msg Msg

	for {
		_, bytes, err := conn.ReadMessage()
		if len(string(bytes)) != 0 {
			if err != nil {
				log.Printf("errror: %s\n", err.Error())
				return
			}

		}

		err = json.Unmarshal(bytes, &msg)
		if err != nil {
			log.Printf("error: %s\n", err.Error())
			return
		}

		fmt.Printf("%s: %s\n", msg.Name, msg.Msg)
	}
}
