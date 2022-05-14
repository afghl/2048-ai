package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"net/http"
)

var (
	addr = flag.String("addr", ":8080", "http service address")
)

func main() {
	fmt.Println("start")
	flag.Parse()
	http.HandleFunc("/ai", ai)
	static := http.FileServer(http.Dir("./2048"))
	http.Handle("/js/", static)
	http.Handle("/style/", static)
	http.Handle("/meta/", static)
	http.Handle("/favicon.ico", static)

	indexTpl := template.Must(template.ParseFiles("./2048/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		indexTpl.Execute(w, nil)
	})

	log.Printf("Service started on \x1b[32;1m%s\x1b[32;1m\x1b[0m\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ai(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("in ai method \n")
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer ws.Close()
	for {
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			break
		}
		fmt.Printf("read message from wb, %v, %s \n", messageType, string(p))
		// write back
		str := string(p)
		if err := ws.WriteMessage(websocket.TextMessage, []byte("echo: "+str)); err != nil {
			fmt.Errorf("write message error: %v \n", err)
		}
	}
}
