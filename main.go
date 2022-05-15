package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/afghl/2048-ai/lib"
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"net/http"
	"time"
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

	html := template.Must(template.ParseFiles("./2048/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_ = html.Execute(w, nil)
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

type req struct {
	Size int     `json:"size"`
	Grid [][]int `json:"grid"`
}

type rsp struct {
	Act int `json:"act"`
}

func ai(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("in ai method \n")
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer ws.Close()
	agent := lib.NewAgent()
	for i := 0; i < 100000; i++ {
		_, p, err := ws.ReadMessage()
		timer := time.NewTimer(300 * time.Millisecond)
		if err != nil {
			break
		}
		m := &req{}
		if err := json.Unmarshal(p, m); err != nil {
			fmt.Errorf("unmarshal error, %v \n", err)
			return
		}
		state := lib.NewState(m.Grid)
		act := agent.GetAction(state)
		r, _ := json.Marshal(rsp{Act: int(act)})
		<-timer.C
		if err := ws.WriteMessage(websocket.TextMessage, r); err != nil {
			fmt.Errorf("write req error: %v \n", err)
		}
	}
}
