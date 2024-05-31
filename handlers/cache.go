package handlers

import (
	apicode "trainee3/lib/apiCode"
	"trainee3/model/output"
	"trainee3/model/output/retStatus"
	"trainee3/services"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ICache interface {
	Refresh(c *fiber.Ctx) error
}

type Cache struct {
	services *services.Services
}

func NewCache(services *services.Services) ICache {
	return &Cache{
		services: services,
	}
}

func (h *Cache) Refresh(c *fiber.Ctx) error {
	if code := h.services.Refresh.ServicesRefresh(); code != apicode.Success {
		logrus.Error(output.CommonResponse{
			RetStatus: retStatus.New(code),
		})
		return c.Status(fiber.StatusOK).JSON(output.CommonResponse{
			RetStatus: retStatus.New(code),
		})
	}
	logrus.Info(output.CommonResponse{
		RetStatus: retStatus.New(apicode.Success),
	})
	return c.Status(fiber.StatusOK).JSON(output.CommonResponse{
		RetStatus: retStatus.New(apicode.Success),
	})
}
