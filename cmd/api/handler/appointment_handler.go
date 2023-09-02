package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/somkieatW/candidate/pkg/core/registry/core"
	"github.com/somkieatW/candidate/pkg/modules/appointment"
	"github.com/somkieatW/candidate/pkg/modules/appointment/models"
	"strconv"
)

type AppointmentAPIHandler struct {
	router             fiber.Router
	CoreRegistry       *core.CoreRegistry
	appointmentUseCase appointment.AppointmentUseCase
}

func NewAppointmentAPIHandler(router fiber.Router, coreRegistry *core.CoreRegistry, appointmentUseCase appointment.AppointmentUseCase) *AppointmentAPIHandler {
	return &AppointmentAPIHandler{
		router:             router,
		CoreRegistry:       coreRegistry,
		appointmentUseCase: appointmentUseCase,
	}
}

func (h *AppointmentAPIHandler) Init() {
	router := h.router

	router.Get("/appointment/list", h.List)
	router.Get("/appointment/{id}", h.Info)
}

//http://localhost:3001/appointment/list?offset=1&pageSize=10

func (h *AppointmentAPIHandler) List(c *fiber.Ctx) error {
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	offset, err := strconv.Atoi(c.Query("offset"))
	request := &models.AppointmentListRequest{PageSize: pageSize, Offset: offset}

	ctx := c.UserContext()
	data, err := h.appointmentUseCase.List(ctx, request)
	if err != nil {
		return err
	}

	c.Status(200).JSON(data)
	return nil
}

func (h *AppointmentAPIHandler) Info(c *fiber.Ctx) error {
	request := &models.AppointmentInfoRequest{ID: c.Params("id")}

	ctx := c.UserContext()
	data, err := h.appointmentUseCase.Info(ctx, request)
	if err != nil {
		return err
	}

	c.Status(200).JSON(data)
	return nil
}
