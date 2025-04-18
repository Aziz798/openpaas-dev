package server

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/joho/godotenv/autoload"
	"openpaas.tech/internal/auth"
)

func (s *FiberServer) RegisterFiberRoutes() {
	origins := os.Getenv("CORS_ORIGINS")
	if origins == "" {
		panic("CORS_ORIGINS environment variable is not set")
	}
	api := s.App.Group("api/v1/")
	api.Use(logger.New(logger.Config{
		Format: "\n[${time}] | [${status}] | [${method}] ${path}\n" +
			"Received: ${bytesReceived} bytes | Sent: ${bytesSent} bytes | " +
			"Latency: ${latency} | IP: ${ip} | Error: ${error}\n",
	}))
	api.Use(recover.New(recover.ConfigDefault))
	api.Use(helmet.New(helmet.ConfigDefault))
	api.Get("/metrics", monitor.New(monitor.Config{Title: "openpaas.tech Metrics"}))
	api.Use(cors.New(cors.Config{
		AllowOrigins:     origins,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Refresh-Token, X-RateLimit-Limit, X-RateLimit-Remaining, X-RateLimit-Reset, Idempotency-Key,X-Cache",
		AllowMethods:     "GET,POST,OPTIONS",
		AllowCredentials: true,
	}))
	api.Use(idempotency.New(idempotency.ConfigDefault))
	api.Use(limiter.New(limiter.Config{
		Max:        10,
		Expiration: 2 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests",
			})
		},
		SkipFailedRequests:     false,
		SkipSuccessfulRequests: false,
	}))
	api.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	api.Get("/health", s.healthHandler)
	auth.RegisterAuthRoutes(api, s.db)

}

func (s *FiberServer) healthHandler(c *fiber.Ctx) error {
	return c.JSON(s.db.Health())
}
