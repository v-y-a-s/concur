package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type user struct {
	Name        string
	Age         rune
	Email       string
	Username    string
	Preferences preference
}

type preference struct {
	EmailAlerts bool
	DisaplyMode string
}

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {

		vyas := user{
			Email: "test@gmail.com",
		}

		c.Response().SetBody(marshall(vyas))
		c.SendStatus(fiber.DefaultConcurrency)
		return c.SendStatus(fiber.StatusAccepted)
	})
	app.Listen(":3000")
}

func marshall(input interface{}) []byte {
	responseBody, _ := json.Marshal(input)
	return responseBody
}
