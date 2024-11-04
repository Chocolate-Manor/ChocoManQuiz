package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var quizCollection *mongo.Collection

func main() {
	app := fiber.New()
	// allow any remote server to make cors requests (not good, but only for development)
	app.Use(cors.New())

	setupDB()
	setupWebsockets(app)

	app.Get("/", index)
	app.Get("/api/quizzes", getQuizzes)

	log.Fatal(app.Listen(":3000"))
}

func setupWebsockets(app *fiber.App) {

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		var (
			mt  int
			msg []byte
			err error
		)
		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", msg)

			if err = c.WriteMessage(mt, msg); err != nil {
				log.Println("write:", err)
				break
			}
		}

	}))
	// Access the websocket server: ws://localhost:3000/ws/123?v=1.0
	// https://www.websocket.org/echo.html
}

func setupDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	quizCollection = client.Database("chocomanQuiz").Collection("quizzes")

	fmt.Println("Connected to MongoDB!")
}

func getQuizzes(c *fiber.Ctx) error {

	// empty map as filter, since we want everything and don't need a filter
	cursor, err := quizCollection.Find(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}

	// a slice of maps with string as key and type of any
	quizzes := []map[string]any{}
	// can decode directly into a Go data type
	err = cursor.All(context.Background(), &quizzes)
	if err != nil {
		panic(err)
	}
	return c.JSON(quizzes)
}

func index(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
