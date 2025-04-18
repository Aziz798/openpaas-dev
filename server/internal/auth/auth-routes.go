package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"openpaas.tech/internal/database"
	"openpaas.tech/internal/types"
	"openpaas.tech/internal/validations"
)

func RegisterAuthRoutes(app fiber.Router, dbInstance database.Service) {
	DB := dbInstance.DB()
	v := validations.GetGlobalValidator()
	api := app.Group("auth/")

	api.Post("signup/", func(c *fiber.Ctx) error {
		return signupWithEmailRouteHandler(c, DB, v)
	})
}

func signupWithEmailRouteHandler(c *fiber.Ctx, db *sqlx.DB, v validations.Validator) error {
	var userSignup types.UserSignupWithEmail
	if err := c.BodyParser(&userSignup); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	// Validate the request
	errList := v.Validate(&userSignup)
	if len(errList) != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "failed", "message": errList})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "message": "User signed up successfully"})
}
