package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func BoardCastMsg() {
	fmt.Println("LOG: Started BoardCasting Messages")

	for {
		msg := <-BoardCast

		for c := range Clients {
			if c != *Sndr {
				err := c.ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`
						{
							"name": "%s", 
							"msg" : "%s"
							}
					`, Sndr.name, msg)))

				if err != nil {
					log.Printf("err: %s\n", err.Error())
					c.ws.Close()
					fmt.Println("Disconnected:", c.name)
					delete(Clients, c)
				}
			}
		}

		Sndr = nil
	}
}
