package service

import (
	"fmt"
	"strings"

	"github.com/gofiber/contrib/websocket"
)

// QuizService as dependency for easy access to quizzes
type NetService struct {
	quizService *QuizService

	// the host client
	host *websocket.Conn

	tick int
}

func Net(quizService *QuizService) *NetService {
	return &NetService{
		quizService: quizService,
	}
}

func (c *NetService) OnIncomingMessage(con *websocket.Conn, mt int, msg []byte) {
	str := string(msg)
	parts := strings.Split(str, ":")

	cmd := parts[0]
	argument := parts[1]

	switch cmd {
	case "host":
		{
			fmt.Println("host quiz", argument)
			c.host = con
			c.tick = 100
			// go routine
			// go func() {
			// 	for {
			// 		c.tick--
			// 		c.host.WriteMessage(websocket.TextMessage, []byte(strconv.Itoa(c.tick)))
			// 		time.Sleep(time.Second) //sleep 1s
			// 	}
			// }()
			break
		}
	case "join":
		{
			fmt.Println("join code", argument)
			// a message visible only to client (frontend)
			c.host.WriteMessage(websocket.TextMessage, []byte("A player joined!"))
			break
		}
	}
}
