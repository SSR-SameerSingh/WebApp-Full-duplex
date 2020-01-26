package main

// var ws = new WebSocket("ws://localhost:8000/ws"), 
// ws.addEventListener("message", function(e) {console.log(e);});, 
// ws.send("bar")

import (
	// "fmt"
	// "log"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{} // upgrades our original http request into websocket

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/v1/ws", func(w http.ResponseWriter, r *http.Request) {
		var conn, _ = upgrader.Upgrade(w, r, nil)
		go func(conn *websocket.Conn) { //important to create for multiple requests
			for {
				mType, msg, _ := conn.ReadMessage() // read message coming from client

				conn.WriteMessage(mType, msg)
			}
		}(conn)
	})

	http.HandleFunc("/v2/ws",func(w http.ResponseWriter, r *http.Request) {
		var conn, _ = upgrader.Upgrade(w, r, nil) // to read/ receive from a client
		go func(conn *websocket.Conn) {
			for {
				_, msg, _ := conn.ReadMessage()
				println(string(msg))
			}
		}(conn)
	})

	http.HandleFunc("/v4/ws",func(w http.ResponseWriter, r *http.Request) {
		var conn, _ = upgrader.Upgrade(w, r, nil) // to read/ receive from a client

		go func(conn *websocket.Conn) {
			for {
				_, _, err := conn.ReadMessage()
				if err != nil {
					conn.Close()
				}
			}
		}(conn)

		go func(conn *websocket.Conn) {
			ch := time.Tick(5 *time.Second)

			for range ch {
				conn.WriteJSON(myStruct{
					Username:"sameer",
					FirstName:"Sameer Singh",
					LastName:"Rathor",
				})
			}
		}(conn)
	})

	http.ListenAndServe(":3000", nil)
}

type myStruct struct {
	Username string `json:"Username"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastname"`
}

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Home Page")
// }

// func wsEndpoint(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "WebSockets")
// }

// func setupRoutes() {
// 	http.HandleFunc("/", homePage)
// 	http.HandleFunc("/ws", wsEndpoint)
// }

// func main() {
// 	fmt.Println("Go WebSockets")
// 	setupRoutes()
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }
