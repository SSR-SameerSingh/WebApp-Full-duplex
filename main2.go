package main

// var ws = new WebSocket("ws://localhost:8000/ws"), 
// ws.addEventListener("message", function(e) {console.log(e);});, 
// ws.send("bar")

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/websocket"

)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(string(p))

		//echo that message into client using conn.WriteMessage
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
		}
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Client Sucessfully Connected...")

	reader(ws)
	// fmt.Fprintf(w, "WebSocket Endpoint")
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("Go WebSockets")
	setupRoutes()

	log.Fatal(http.ListenAndServe(":8080", nil))
}