package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func ResigerAPP(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

}
func User(app *fiber.App) {
	api := app.Group("/v1.0")
	group := api.Group("/user")
	group.Post("/webhook_log", facebook.WebhookLog)
}
