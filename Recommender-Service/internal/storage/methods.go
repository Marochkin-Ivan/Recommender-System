package storage

import (
	"github.com/gofiber/fiber/v2"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

func (c *Config) Init() {
	file, err := os.Open("./configs/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	d := yaml.NewDecoder(file)

	if err := d.Decode(&c); err != nil {
		log.Fatal(err)
	}
}

func (s *Server) Init(c Config, u Storable, t Storable, a *fiber.App) {
	s.UserStore = u
	s.TicketsStore = t
	s.Cfg = c
	s.App = a
}
