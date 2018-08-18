package main

import (
	"time"

	"github.com/gorilla/websocket"
)

// client represents a single chatting user .
// socket is the websocket for this client.
// send is a channel on which messages are sent.
// room is a room this client is chatting in.
type client struct {
	socket *websocket.Conn
	send   chan *message
	room   *room
	//userData holds information about the user
	userData map[string]interface{}
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		var msg *message
		err := c.socket.ReadJSON(&msg)
		if err != nil {
			return
		}
		msg.When = time.Now()
		msg.Name = c.userData["name"].(string)
		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteJSON(msg)
		if err != nil {
			break
		}
	}
}
