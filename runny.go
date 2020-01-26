package main 
// var ws = new WebSocket("ws://localhost:8000/ws"), 
// ws.addEventListener("message", function(e) {console.log(e);});, 
// ws.send("bar")

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/websocket"
	"sync"
)

var wg sync.WaitGroup
var upgrader = websocket.Upgrader{}
// 	ReadBufferSize: 1024,
// 	WriteBufferSize: 1024,
// }

func hello(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "fuck You")
	http.ServeFile(w, r, "index2.html")
}

func ws(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		// fmt.Fprintf(w, "Error in establishing connection")
		log.Println(string("Checking...."))
		log.Fatal(err)
	}
	// defer conn.Close()
	// msg := []byte("Lets start to talk something.")
	// err = conn.WriteMessage(websocket.TextMessage, msg)
	// if err != nil {
	// 	log.Println(err)
	// }
	// wg.Add(1)
	func(conn *websocket.Conn) { //reader function
		// defer wg.Done()
		for {
			mType, msg, err := conn.ReadMessage()
			if err != nil {
				log.Printf("Failed to read message %v", err)
				conn.Close()
				return
			}
			// log.Println(string("yolo"))
			log.Println(string(msg)) // print incoming message
			if err := conn.WriteMessage(mType, msg); err != nil { //echo the message back to client
				log.Println(err)
				return
			} 
		}
	}(conn)
	// wg.Wait()
	// conn.Close()

	// go func(conn *websocket.Conn) {
	// 	conn.WriteMessage(mtype, msg)
	// }(conn)

	// for {
	// 	mType, msg, err := conn.ReadMessage()
	// 	if err != nil {
	// 		log.Printf("failed to read message %v", err)
	// 		conn.Close()
	// 		return
	// 	}
	// 	log.Println(string(msg))
	// 	conn.WriteMessage(mType, msg)
	// }
	fmt.Fprintf(w, "Connection Established sucessfully")
	fmt.Fprintf(w, "YAY !")
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/ws", ws)
	http.ListenAndServe(":8080", nil)
}