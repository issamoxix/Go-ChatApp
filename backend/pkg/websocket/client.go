package websocket

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID    string
	Conn  *websocket.Conn
	Pool  *Pool
	mutex sync.Mutex
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		fmt.Println("Closing Connection Inside Read")
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()
	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		message := Message{
			Type: messageType,
			Body: string(p),
		}
		c.Pool.Broadcast <- message
		fmt.Printf("Message received: %s", message)

	}
}
