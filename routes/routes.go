package routes

import (
    "github.com/liqlvnvn/go-auth/controllers"
    "github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
    app.Get("/", controllers.Hello)
}
