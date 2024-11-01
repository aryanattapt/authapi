package controller

import (
	"authapi/model"
	"authapi/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func SignUp(ctx *fiber.Ctx) error {
	if ctx.Method() == "POST" {
		var payload = &model.UserInsertPayload{}
		if err := ctx.BodyParser(payload); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code":    "SIGNUP.INVALIDPAYLOAD.EXCEPTION",
				"message": "Sorry, System can't parse your data! Please Recheck!",
				"error":   err.Error(),
			})
		}

		var goValidator = validator.New()
		if err := goValidator.Struct(payload); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code":    "SIGNUP.INVALIDPAYLOAD.EXCEPTION",
				"message": "Parameter is required!",
				"error":   err.Error(),
			})
		}

		err := service.SignUp(payload)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":  "SIGNUP.UNAUTHORIZED.EXCEPTION",
				"error": err.Error(),
			})
		}

		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Success Signup!",
		})
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}

func SignIn(ctx *fiber.Ctx) error {
	if ctx.Method() == "POST" {
		var payload = &model.SignInPayload{}
		if err := ctx.BodyParser(payload); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code":    "SIGNIN.INVALIDPAYLOAD.EXCEPTION",
				"message": "Sorry, System can't parse your data! Please Recheck!",
				"error":   err.Error(),
			})
		}

		var goValidator = validator.New()
		if err := goValidator.Struct(payload); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code":    "SIGNIN.INVALIDPAYLOAD.EXCEPTION",
				"message": "Parameter is required!",
				"error":   err.Error(),
			})
		}

		err, token := service.Signin(payload)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":  "SIGNUP.UNAUTHORIZED.EXCEPTION",
				"error": err.Error(),
			})
		}

		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Success Signin!",
			"token":   token,
		})
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}
