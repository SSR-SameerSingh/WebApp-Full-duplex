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
	// "time"
)

var wg sync.WaitGroup
var upgrader = websocket.Upgrader{}
// var upgrader2 = websocket.Upgrader{}

type longLatStruct struct {
	Long float64 `json:"longitute"`
	Lat float64 `jason:"latitude"`
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan *longLatStruct)

func index1(w http.ResponseWriter, r *http.Request) {
	log.Println("index1 fun")
	http.ServeFile(w, r, "assets/index1.html")
}

func index2(w http.ResponseWriter, r *http.Request) {
	log.Println("index2 fun")
	http.ServeFile(w, r, "assets/index2.html")
}

func index3(w http.ResponseWriter, r *http.Request) {
	log.Println("index3 fun")
	http.ServeFile(w, r, "assets/index3.html")
}

func index4(w http.ResponseWriter, r *http.Request) {
	log.Println("index4 fun")
	http.ServeFile(w, r, "assets/index4.html")
}

func ws(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	log.Println("connection upgraded fun")
	clients[conn] = true
	if err != nil {
		log.Println(string("Checking...."))
		log.Fatal(err)
	}

	// go func(conn *websocket.Conn) {
	// 		ch := time.Tick(100 *time.Millisecond)
	// 		// counter := 0

	// 		for range ch {
	// 			mType, msg, err := conn.ReadMessage()
	// 			if err != nil {
	// 				log.Printf("Failed to read message, %v", err)
	// 			}
	// 			for client := range clients {
	// 				err := client.WriteMessage(websocket.TextMessage, msg)
	// 				if err != nil {
	// 					log.Printf("Error sending message: %v", err)
	// 					client.Close()
	// 					delete(clients, client)
	// 				}
	// 			}
	// 			conn.WriteMessage(mType, msg)
	// 			// msg := []byte(string(counter))
	// 			// conn.WriteMessage(websocket.TextMessage, msg)
	// 			// counter += 1
	// 		}
	// 	}(conn)
	func(conn *websocket.Conn) { //reader function
		for {
			mType, msg, err := conn.ReadMessage()
			if err != nil {
				log.Printf("Failed to read message %v", err)
				conn.Close()
				return
			}
			for client := range clients {
					err := client.WriteMessage(mType, msg)
					if err != nil {
						log.Printf("Error sending message: %v", err)
						client.Close()
						delete(clients, client)
					}
				}
			log.Println(string(msg)) // print incoming message
			// if err := conn.WriteMessage(mType, msg); err != nil { //echo the message back to client
			// 	log.Println(err)
			// 	return
			// } 
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