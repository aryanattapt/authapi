package router

import (
	"authapi/controller"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app *fiber.App) {
	authRouter := app.Group("/auth")
	authRouter.All("/signin", controller.SignIn)
	authRouter.All("/signup", controller.SignUp)
}
