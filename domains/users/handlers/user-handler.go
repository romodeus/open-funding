package userhandler

import (
	"errors"
	"net/http"
	entity "open-funding/domains/users/entities"
	"open-funding/exceptions"
	"open-funding/utils/helpers"
	"strconv"

	"open-funding/validation"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	Usecase entity.IusecaseUser
}

func New(usecase entity.IusecaseUser) *userHandler {
	return &userHandler{
		Usecase: usecase,
	}
}

func (h *userHandler) Store(c *fiber.Ctx) error {
	request := Request{}
	err := c.BodyParser(&request)
	if err != nil {
		return errors.New(exceptions.BadRequestErrorMsg)
	}

	errResponse := validation.ValidateStruct(request)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(helpers.FailedResponseData("bad request", errResponse))
	}

	h.Usecase.Store(RequestToEntity(request))

	return c.JSON(helpers.SuccessGetResponse("success create data"))
}

func (h *userHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	request := Request{}
	if err != nil {
		return errors.New(exceptions.BadRequestErrorMsg)
	}

	err = c.BodyParser(&request)
	if err != nil {
		return errors.New(exceptions.BadRequestErrorMsg)
	}

	errResponse := validation.ValidateStruct(request)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(helpers.FailedResponseData("bad request", errResponse))
	}

	userEntity := RequestToEntity(request)
	userEntity.UID = uint(id)

	h.Usecase.Update(userEntity)

	return c.Status(http.StatusOK).JSON(helpers.SuccessGetResponse("success update data"))
}
