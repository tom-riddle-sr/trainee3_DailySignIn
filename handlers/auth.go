package handlers

import (
	"trainee3/lib/apiCode"
	"trainee3/model/input"
	"trainee3/model/output"
	"trainee3/model/output/retStatus"
	"trainee3/services"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ISignIn interface {
	SignIn(c *fiber.Ctx) error
}

type SignIn struct {
	services *services.Services
}

func NewActivity(services *services.Services) ISignIn {
	return &SignIn{
		services: services,
	}
}

func (s *SignIn) SignIn(c *fiber.Ctx) error {
	inputData := input.MemberIdRequest{}
	if err := c.BodyParser(&inputData); err != nil {
		logrus.Error(output.CommonResponse{
			RetStatus: retStatus.New(apiCode.ValidateError),
		})
		return c.Status(fiber.StatusOK).JSON(output.CommonResponse{
			RetStatus: retStatus.New(apiCode.ValidateError),
		})
	}

	validate := validator.New()
	if err := validate.Struct(inputData); err != nil {
		logrus.Error("SignIn validate error:", output.CommonResponse{
			RetStatus: retStatus.New(apiCode.ValidateError)})
		return c.Status(fiber.StatusOK).JSON(
			output.CommonResponse{
				RetStatus: retStatus.New(apiCode.ValidateError),
			},
		)
	}

	res, code := s.services.Activity.SignIn(inputData)
	if code != apiCode.Success {
		logrus.Error("SignIn error:", output.CommonResponse{
			RetStatus: retStatus.New(code),
		})
		return c.Status(fiber.StatusOK).JSON(output.CommonResponse{
			RetStatus: retStatus.New(code),
		})
	}
	logrus.Info(output.CommonResponse{
		RetStatus: retStatus.New(apiCode.Success),
		Data:      res,
	})
	return c.Status(fiber.StatusOK).JSON(output.CommonResponse{
		RetStatus: retStatus.New(apiCode.Success),
		Data:      res,
	})
}
