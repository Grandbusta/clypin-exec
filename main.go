package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gorilla/websocket"
	"golang.design/x/clipboard"
)

func init() {
	err := clipboard.Init()
	fmt.Println("Clypin running...")
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:3000/ws/send/123", nil)
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
	}

}

func watcher() {
	_, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	env := os.Environ()
	fmt.Println(env[0], env[1])
	changed := clipboard.Watch(context.Background(), clipboard.FmtText)
	for data := range changed {

		println(string(data))
		// if ent, ok := sockets["123"]; ok {
		// 	ent.msgs = append(sockets["123"].msgs, string(data))
		// 	sockets["123"] = ent
		// 	curr := sockets["123"]
		// 	var (
		// 		msg string
		// 		err error
		// 	)
		// 	msg = curr.msgs[len(curr.msgs)-1]
		// 	for _, soc := range curr.conn {
		// 		err = soc.WriteJSON(map[string]string{"msg": msg})
		// 		if err != nil {
		// 			// log.Println("write:", err)
		// 			break
		// 		}
		// 	}
		// }
	}
}

func main() {

}
