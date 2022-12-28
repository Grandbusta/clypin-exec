package internal

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"golang.design/x/clipboard"
)

func InitializeClipboard() {
	err := clipboard.Init()
	fmt.Println("Clypin running...")
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectToServer(url string) (conn *websocket.Conn) {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Can't connect to clypin server")
	}
	log.Println("Connected to clypin...")
	return conn
}
