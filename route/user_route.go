package route

import (
    "github.com/gofiber/fiber/v2"
    "github.com/programmerug/fibermongo/controller"
)

func UserRoute(app *fiber.App) {
    app.Post("/user", controller.CreateUser)
}