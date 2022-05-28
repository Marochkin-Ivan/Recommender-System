package algo

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"log"
	"recommenderalgo/internal/algoStorage"
	"recommenderalgo/internal/db"
)

func StartService() {
	var (
		cfg algoStorage.Config
		d   db.DB
		s   algoStorage.AlgoServer
	)

	// _______start INIT_______
	cfg.Init()
	d.Init(cfg.Store)

	s.Init(cfg, &d, fiber.New())
	s.App.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	s.App.Use(recover.New(recover.Config{EnableStackTrace: true}))
	s.App.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path} ${resBody}\n",
	}))
	// _______end INIT_______

	// ____start HANDLERS____

	// documentation
	s.App.Get("/swagger/*", fiberSwagger.WrapHandler)
	s.App.Post("/", s.Recommend)
	// ____end HANDLERS____

	chanIn := make(chan algoStorage.RecommendTuple, 100)
	s.AlgoChan = chanIn
	go s.AlgoSolve(s.AlgoChan)
	log.Println(s.App.Listen(s.Cfg.AlgHost))

	close(chanIn)
}
