package routes

import (
    "github.com/gofiber/fiber/v2"
    "payments/src/controller" // replace
)

func CatchphrasesRoute(route fiber.Router) {
    route.Get("/", controllers.GetAllCatchphrases)
    route.Get("/:id", controllers.GetCatchphrase)
    route.Post("/add", controllers.Addphrases)
    route.Put("/:id", controllers.UpdateCatchphrase)
    route.Delete("/:id", controllers.DeleteCatchphrase)
}