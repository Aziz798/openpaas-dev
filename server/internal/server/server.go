package server

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"openpaas.tech/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "openpaas.tech",
			AppName:      "openpaas.tech",
			JSONEncoder:  json.Marshal,
			JSONDecoder:  json.Unmarshal,
		}),

		db: database.New(),
	}

	return server
}
