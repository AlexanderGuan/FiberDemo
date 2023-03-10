package routes

import (
	"github.com/AlexanderGuan/FiberDemo/app/controllers"
	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for GET method:

	// Routes for POST method:
	route.Post("/user/sign/up", controllers.UserSignUp) // register a new user
	// route.Post("/user/sign/in", controllers.UserSignIn) // auth, return Access & Refresh tokens
}
