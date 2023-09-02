package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/somkieatW/interview-appointment/pkg/core/registry/core"
	models2 "github.com/somkieatW/interview-appointment/pkg/models"
	"github.com/somkieatW/interview-appointment/pkg/modules/appointment"
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
	router.Get("/appointment/:id", h.Info)
}

func (h *AppointmentAPIHandler) List(c *fiber.Ctx) error {
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	offset, err := strconv.Atoi(c.Query("offset"))
	request := &models2.AppointmentListRequest{PageSize: pageSize, Offset: offset}

	ctx := c.UserContext()
	data, err := h.appointmentUseCase.List(ctx, request)
	if err != nil {
		return err
	}

	c.Status(200).JSON(data)
	return nil
}

func (h *AppointmentAPIHandler) Info(c *fiber.Ctx) error {
	request := &models2.AppointmentInfoRequest{ID: c.Params("id")}

	ctx := c.UserContext()
	data, err := h.appointmentUseCase.Info(ctx, request)
	if err != nil {
		return err
	}

	c.Status(200).JSON(data)
	return nil
}
