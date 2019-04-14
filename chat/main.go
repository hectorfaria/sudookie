package main

import (
	"fmt"
	"net/http"

	"./pkg/websocket"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Enter Da WebSocket (ﾉ◕ヮ◕)ﾉ*:・ﾟ✧ ")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {
	fmt.Println("Server Started on port 8080 ლ( ̅°̅ ੪ ̅°̅ )ლ Bæsj")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
