package main

import (
	"flag"
	"fmt"
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
	http.HandleFunc("/ai", ai)

	log.Printf("Service started on \x1b[32;1m%s\x1b[32;1m\x1b[0m\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func ai(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("in ai method %v, %v\n", w, r)
}
