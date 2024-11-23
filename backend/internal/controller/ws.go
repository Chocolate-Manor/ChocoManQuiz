package controller

import (
	"log"

	"chocomanquiz.com/quiz/internal/service"
	"github.com/gofiber/contrib/websocket"
)

type WebsocketController struct {
	netService *service.NetService
}

func Ws(netService *service.NetService) WebsocketController {
	return WebsocketController{
		netService: netService,
	}
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

		c.netService.OnIncomingMessage(con, mt, msg)

		if err = con.WriteMessage(mt, msg); err != nil {
			log.Println("write:", err)
			break
		}
	}
}
