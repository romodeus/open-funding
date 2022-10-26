package exceptions

import (
	"net/http"
	"open-funding/utils/helpers"

	"github.com/gofiber/fiber/v2"
)

var (
	InternalServerErrorMsg = "internal server error"
	ForbiddenErrorMsg      = "forbidden error"
	NotFoundErrorMsg       = "not found error"
	BadRequestErrorMsg     = "bad request error"
	ValidationErrorMsg     = "validation error"
)

func ExceptionHandler(ctx *fiber.Ctx, err error) error {
	if NotFoundError(err, ctx) {
		return nil
	} else if ForbiddenError(err, ctx) {
		return nil
	} else if BadRequestError(err, ctx) {
		return nil
	} else {
		InternalServerError(err, ctx)
		return nil
	}
}

func InternalServerError(err error, c *fiber.Ctx) bool {

	if err.Error() == InternalServerErrorMsg {
		c.Status(http.StatusInternalServerError).JSON(helpers.FailedResponse(err.Error()))
		return true
	}

	return false
}

func NotFoundError(err error, c *fiber.Ctx) bool {

	if err.Error() == NotFoundErrorMsg {
		c.Status(http.StatusNotFound).JSON(helpers.FailedResponse(err.Error()))
		return true
	}

	return false
}

func ForbiddenError(err error, c *fiber.Ctx) bool {

	if err.Error() == ForbiddenErrorMsg {
		c.Status(http.StatusForbidden).JSON(helpers.FailedResponse(err.Error()))
		return true
	}

	return false
}

func BadRequestError(err error, c *fiber.Ctx) bool {

	if err.Error() == BadRequestErrorMsg {
		c.Status(http.StatusBadRequest).JSON(helpers.FailedResponse(err.Error()))
		return true
	}

	return false
}
