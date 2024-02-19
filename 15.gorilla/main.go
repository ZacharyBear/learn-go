package main

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/gorilla/websocket"
)

func main() {
	wsUpgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// Upgrade HTTP connection to a WebSocket connection.
		conn, err := wsUpgrader.Upgrade(w, r, nil)
		if err != nil {
			color.Red(err.Error())
			return
		}
		// Read messages
		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				color.Red(err.Error())
				return
			}

			// Print the message to the console
			color.Cyan("Received: %s", message)

			// Send a message back to the client
			err = conn.WriteMessage(messageType, []byte(fmt.Sprintf("We've got your message: %s", message)))
			if err != nil {
				color.Red(err.Error())
				return
			}
		}
	})

	http.ListenAndServe(":888", nil)
}
