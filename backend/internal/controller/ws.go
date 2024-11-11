package controller

import (
	"log"

	"github.com/gofiber/contrib/websocket"
)

type WebsocketController struct {
}

func Ws() WebsocketController {
	return WebsocketController{}
}

func (c WebsocketController) Ws(con *websocket.Conn) {
	var mt int
	var msg []byte
	var err error
	for {
		if mt, msg, err = con.ReadMessage(); err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", msg)

		if err = con.WriteMessage(mt, msg); err != nil {
			log.Println("write:", err)
			break
		}
	}
}
