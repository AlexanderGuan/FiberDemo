package main

import (
	"os"

	"github.com/AlexanderGuan/FiberDemo/pkg/routes"
	"github.com/AlexanderGuan/FiberDemo/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Define a new Fiber app with config.
	app := fiber.New()

	// Routes.
	routes.PublicRoutes(app)

	// Start server (with or without a graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
