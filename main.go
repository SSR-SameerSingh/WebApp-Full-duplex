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
	"time"
)

type myStruct struct {
	Username string `json:"Username"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastname"`
}

var wg sync.WaitGroup
var upgrader = websocket.Upgrader{}

func index1(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index2.html")
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
		// fmt.Fprintf(w, "Error in establishing connection")
		log.Println(string("Checking...."))
		log.Fatal(err)
	}

	go func(conn *websocket.Conn) {
			ch := time.Tick(5 *time.Second)
			counter := 0

			for range ch {
				msg := []byte(string(counter))
				conn.WriteMessage(websocket.TextMessage, msg)
				counter += 1
				// conn.WriteJSON(myStruct{
				// 	Username:"ssr.sameersingh",
				// 	FirstName:"Sameer Singh",
				// 	LastName:"Rathor",
				// })
			}
		}(conn)

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
	log.Println("Server Started.....")
	http.HandleFunc("/index1", index1)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/index3", index3)
	http.HandleFunc("/index4", index4)
	http.HandleFunc("/ws", ws)
	http.ListenAndServe(":8080", nil)
}