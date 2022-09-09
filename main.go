package main

import (
	"first_demo/model/student"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/demo", func(c *fiber.Ctx) error {
		var body student.Student
		err := c.BodyParser(body)
		fmt.Println("====", body, "+++", err)
		return c.JSON(
			map[string]interface{}{
				"data": body,
				"err":  err,
			},
		)
	})

	app.Listen(":3000")
}
