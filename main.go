package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	app := fiber.New(&fiber.Settings{
		CaseSensitive: false,
		StrictRouting: false,
		ServerHeader:  "GW-AS-A-SERVICE",
	})

	app.Use(logger.New())

	app.Static("/", "./resources/site")

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/grunt", func(c *fiber.Ctx) {
		c.Set("Cache-Control", "no-store")
		gruntID := rand.Intn(5) + 1
		c.SendFile("./resources/gw/grunt/" + strconv.Itoa(gruntID) + ".ogg")
	})

	v1.Get("/aaa", func(c *fiber.Ctx) {
		c.Set("Cache-Control", "no-store")
		aaaID := rand.Intn(7) + 1
		c.SendFile("./resources/gw/aaa/" + strconv.Itoa(aaaID) + ".ogg")
	})

	app.Listen(3000)
}
