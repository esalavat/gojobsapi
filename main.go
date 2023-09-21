package main

import (
	"log"

	"github.com/esalavat/gojobsapi/db"
	"github.com/esalavat/gojobsapi/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db.ConnectDb()

	app := fiber.New()

	app.Use(cors.New())

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))

}

func setupRoutes(app *fiber.App) {
	app.Get("/", handleRoot)

	app.Post("/jobs", routes.CreateJob)
	app.Get("/jobs", routes.GetJobs)
	app.Get("/jobs/:id", routes.GetJob)
	app.Put("/jobs/:id", routes.UpdateJob)
	app.Delete("/jobs/:id", routes.DeleteJob)
}

func handleRoot(c *fiber.Ctx) error {

	return c.SendString("Welcome to the API")
}
