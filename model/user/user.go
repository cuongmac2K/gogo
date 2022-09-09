package user

import (
	"first_demo/config"
	"first_demo/helpers"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(ctx *fiber.Ctx) error {
	body := new(User)

	if err := ctx.BodyParser(body); err != nil {
		return ctx.JSON(helpers.Response{
			Status:  false,
			Message: "Missing params.",
			Error: helpers.Error{
				ErrorCode:    config.ErrorCode["ERROR_MISSING_PARAMS"],
				ErrorMessage: err.Error(),
			},
		})
	}
	return nil
}
