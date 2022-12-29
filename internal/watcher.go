package internal

import (
	"clypin-exec/internal/structs"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"golang.design/x/clipboard"
)

func RunWatcher() {
	conn := ConnectToServer("ws://localhost:3000/ws/send/123?token=fksjanjndajndsd9ydbseqoQ")
	defer conn.Close()
	_, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	env := os.Environ()
	fmt.Println(env[0], env[1])
	changed := clipboard.Watch(context.Background(), clipboard.FmtText)
	for data := range changed {
		println(string(data))
		var (
			err error
		)
		err = conn.WriteJSON(map[string]string{
			"msg":    string(data),
			"source": "watcher",
		})
		if err != nil {
			log.Println("write:", err)
			break
		}
		var rcvData structs.MsgData
		if err = conn.ReadJSON(&rcvData); err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", data)
	}

}
