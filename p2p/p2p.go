package p2p

import (
	"fmt"
	"net/http"

	"github.com/GeonHyeok-Lee/minimal-cryptocurrency/utils"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)
	for {
		_, p, err := conn.ReadMessage()
		utils.HandleErr(err)
		fmt.Printf("%s\n\n", p)
	}
}
