package app

import (
	"log"
	"recommendersystem/internal/db"
	"recommendersystem/internal/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func StartService() {
	var (
		cfg storage.Config
		d   db.DB
		s   storage.Server
	)

	// _______start INIT_______
	cfg.Init()
	d.Init(cfg.Store)

	s.Init(cfg, &d, &d, fiber.New())
	s.App.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	s.App.Use(recover.New(recover.Config{EnableStackTrace: true}))
	s.App.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path} ${resBody}\n",
	}))
	// _______end INIT_______

	// groups
	api := s.App.Group("/api")

	user := api.Group("/user")
	leisure := api.Group("/leisure")
	tickets := api.Group("/tickets")

	// ____start HANDLERS____
	// users handlers
	user.Post("/registration", s.Registration)
	user.Post("/auth", s.Login)

	// tickets handlers
	tickets.Post("/", s.Tickets)

	// leisure handlers
	leisure.Post("/list", s.LeisureList) // redirect

	// documentation
	api.Get("/swagger/*", fiberSwagger.WrapHandler)
	// ____end HANDLERS____

	log.Fatal(s.App.Listen(s.Cfg.Host))
}
