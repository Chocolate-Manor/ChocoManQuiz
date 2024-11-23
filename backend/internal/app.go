package internal

import (
	"context"
	"log"
	"time"

	"chocomanquiz.com/quiz/internal/collection"
	"chocomanquiz.com/quiz/internal/controller"
	"chocomanquiz.com/quiz/internal/service"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var quizCollection *mongo.Collection

type App struct {
	httpServer *fiber.App
	database   *mongo.Database

	quizService *service.QuizService
	netService  *service.NetService
}

func (a *App) Init() {
	a.setupDb()
	a.setupServices()
	a.setupHttpAndWs()

	log.Fatal(a.httpServer.Listen(":3000"))
}

func (a *App) setupHttpAndWs() {
	hs := fiber.New()
	hs.Use(cors.New())

	quizController := controller.Quiz(a.quizService)
	hs.Get("/api/quizzes", quizController.GetQuizzes)
	hs.Get("/api/quizzes/:quizId", quizController.GetQuizById)
	//hs.Put("/api/quizzes/:quizId", quizController.UpdateQuizById)

	// websocket set up
	wsController := controller.Ws(a.netService)
	hs.Get("/ws", websocket.New(wsController.Ws))

	a.httpServer = hs
}

func (a *App) setupServices() {
	a.quizService = service.Quiz(collection.Quiz(a.database.Collection("quizzes")))
	a.netService = service.Net(a.quizService)
}

func (a *App) setupDb() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	a.database = client.Database("chocomanQuiz")
}
