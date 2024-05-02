package main

import (
	"fmt"
	"net/http"

	"ChatApp/pkg/websocket"
)

func serveWS(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("webscoket endpoint reached")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Errorf("Error Serving WS %s", err)
		return
	}
	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWS(pool, w, r)
	})
}

func main() {
	var addr string = "9000"
	fmt.Println("Server is Running on", addr)
	setRoutes()
	http.ListenAndServe(":"+addr, nil)
}
