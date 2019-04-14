package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

// Client has an ID, connection, and the pool
type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

// Message has a Type of Int and the Body
type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Print(err)
			return
		}
		message := Message{Type: messageType, Body: string(p)}
		c.Pool.Broadcast <- message
		fmt.Printf("Message Recieved (◕‿◕✿): %+v\n", message)
	}
}
