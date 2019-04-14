package main

import (
	"fmt"
	"net/http"

	"./pkg/websocket"
)

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	go websocket.Writer(ws)
	websocket.Reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Server Started on port 8080 ლ( ̅°̅ ੪ ̅°̅ )ლ Bæsj")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
