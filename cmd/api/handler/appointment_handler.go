package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/somkieatW/interview-appointment/pkg/core/registry/core"
	"github.com/somkieatW/interview-appointment/pkg/core/utils"
	"github.com/somkieatW/interview-appointment/pkg/models"
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
	router.Patch("/appointment/", h.Update)
	router.Put("/appointment/archive", h.Archive)
}

func (h *AppointmentAPIHandler) List(c *fiber.Ctx) error {
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil {
		response := utils.ErrorResponse{
			Message: "Invalid request data",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		response := utils.ErrorResponse{
			Message: "Invalid request data",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	request := &models.AppointmentListRequest{PageSize: pageSize, Offset: offset}

	ctx := c.UserContext()
	data, err := h.appointmentUseCase.List(ctx, request)
	if err != nil {
		response := utils.ErrorResponse{
			Message: "Internal server error",
		}
		log.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(data)
}

func (h *AppointmentAPIHandler) Info(c *fiber.Ctx) error {
	request := &models.AppointmentInfoRequest{ID: c.Params("id")}

	ctx := c.UserContext()
	data, err := h.appointmentUseCase.Info(ctx, request)
	if err != nil {
		response := utils.ErrorResponse{
			Message: "Internal server error",
		}
		log.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(data)
}

func (h *AppointmentAPIHandler) Update(c *fiber.Ctx) error {
	request := &models.AppointmentUpdateRequest{}
	if err := c.BodyParser(request); err != nil {
		response := utils.ErrorResponse{
			Message: "Invalid request data",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	ctx := c.UserContext()
	data, err := h.appointmentUseCase.Update(ctx, request)
	if err != nil {
		response := utils.ErrorResponse{
			Message: "Internal server error",
		}
		log.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(data)
}

func (h *AppointmentAPIHandler) Archive(c *fiber.Ctx) error {
	request := &models.AppointmentArchiveRequest{}
	if err := c.BodyParser(request); err != nil {
		response := utils.ErrorResponse{
			Message: "Invalid request data",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	ctx := c.UserContext()
	data, err := h.appointmentUseCase.Archive(ctx, request)
	if err != nil {
		response := utils.ErrorResponse{
			Message: "Internal server error",
		}
		log.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(data)
}
