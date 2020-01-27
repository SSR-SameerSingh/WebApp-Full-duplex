package main 
// var ws = new WebSocket("ws://localhost:8000/ws"), 
// ws.addEventListener("message", function(e) {console.log(e);});, 
// ws.send("bar")

import (
	"fmt"
	"net/http"
	// "net"
	"log"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

var wg sync.WaitGroup
var upgrader = websocket.Upgrader{}
// var upgrader2 = websocket.Upgrader{}

type myStruct struct {
	Username string `json:"Username"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastname"`
}

func index1(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index1.html")
}

func index2(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index2.html")
}

func index3(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index3.html")
}

func index4(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index4.html")
}

func ws(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(string("Checking...."))
		log.Fatal(err)
	}

	go func(conn *websocket.Conn) {
			ch := time.Tick(100 *time.Millisecond)
			// counter := 0

			for range ch {
				mType, msg, err := conn.ReadMessage()
				if err != nil {
					log.Printf("Failed to read message, %v", err)
				}

				conn.WriteMessage(mType, msg)
				// msg := []byte(string(counter))
				// conn.WriteMessage(websocket.TextMessage, msg)
				// counter += 1
			}
		}(conn)
	func(conn *websocket.Conn) { //reader function
		for {
			mType, msg, err := conn.ReadMessage()
			if err != nil {
				log.Printf("Failed to read message %v", err)
				conn.Close()
				return
			}
			log.Println(string(msg)) // print incoming message
			if err := conn.WriteMessage(mType, msg); err != nil { //echo the message back to client
				log.Println(err)
				return
			} 
		}
	}(conn)
	fmt.Fprintf(w, "Connection Established sucessfully")
	fmt.Fprintf(w, "YAY !")
}

func main() {
	log.Println("Server Started.....")
	http.HandleFunc("/index1", index1)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/index3", index3)
	http.HandleFunc("/index4", index4)
	http.HandleFunc("/ws", ws)
	// http.HandleFunc("/send", ss)

	http.ListenAndServe(":8123", nil)
}