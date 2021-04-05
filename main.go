package main

import (
    "github.com/liqlvnvn/go-auth/database"
    "github.com/liqlvnvn/go-auth/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
    database.Connect()

    app := fiber.New()
    routes.Setup(app)

    app.Listen(":3000")
}
