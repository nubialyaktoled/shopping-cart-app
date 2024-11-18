package routes

import (
	"api/controllers"
	"api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/cart/:product_id/:quantity", middlewares.Auth, controllers.AddToCart)
	app.Delete("/cart/:product_id", middlewares.Auth, controllers.RemoveFromCart)
	app.Put("/cart/:product_id/:quantity", middlewares.Auth, controllers.UpdateCartQuantity)
	app.Get("/cart", middlewares.Auth, controllers.GetCart)
	app.Get("/cart/total", middlewares.Auth, controllers.GetCartTotal)
}
