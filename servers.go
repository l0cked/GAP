package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var servers Servers

type Servers struct {
	Users []websocket.Conn
}

func (srv *Servers) Init() {
	http.HandleFunc("/", srv.HttpHandler)
	http.HandleFunc("/ws", srv.WsHandler)

	go http.ListenAndServe(config.BindAddress, nil)
}

func (srv *Servers) HttpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		switch r.URL.Path {
		case "/":
			http.ServeFile(w, r, config.HttpDirectory+"/index.html")
		case "/favicon.ico":
			http.ServeFile(w, r, config.HttpDirectory+"/favicon.ico")
		case "/style.css":
			http.ServeFile(w, r, config.HttpDirectory+"/style.css")
		case "/script.js":
			http.ServeFile(w, r, config.HttpDirectory+"/script.js")
		default:
			http.Error(w, "Not found", http.StatusNotFound)
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (srv *Servers) WsHandler(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	go srv.WsReader(conn)
}

func (srv *Servers) WsReader(conn *websocket.Conn) {
	srv.Register(conn)
	type Message struct {
		Type  string
		Value string
	}
	var msg Message
	for {
		if err := conn.ReadJSON(&msg); err != nil {
			break
		}
		fmt.Println("[WS] Incoming message:", msg)
	}
	srv.Unregister(conn)
}

func (srv *Servers) Register(conn *websocket.Conn) {
	srv.Users = append(srv.Users, *conn)
	srv.SendUsers()
}

func (srv *Servers) Unregister(conn *websocket.Conn) {
	if err := conn.Close(); err != nil {
		panic(err)
	}
	srv.RemoveUser(conn)
	srv.SendUsers()
}

func (srv *Servers) RemoveUser(conn *websocket.Conn) {
	var newUsers []websocket.Conn
	for _, user := range srv.Users {
		if user.RemoteAddr() != conn.RemoteAddr() {
			newUsers = append(newUsers, user)
		}
	}
	srv.Users = newUsers
}

func (srv *Servers) SendUsers() {
	if len(srv.Users) == 0 {
		return
	}
	type Message struct {
		Users int
	}
	for _, conn := range srv.Users {
		if err := conn.WriteJSON(Message{Users: len(srv.Users)}); err != nil {
			panic(err)
		}
	}
}
