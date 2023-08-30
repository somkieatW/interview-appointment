package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/somkieatW/candidate/pkg/core/registry/core"
	"github.com/somkieatW/candidate/pkg/modules/appointment"
	"github.com/somkieatW/candidate/pkg/modules/appointment/models"
	"net/http"
)

type AppointmentAPIHandler struct {
	router             gin.IRouter
	CoreRegistry       *core.CoreRegistry
	AppointmentUseCase appointment.AppointmentUseCase
}

func NewAppointmentAPIHandler(router gin.IRouter, coreRegistry *core.CoreRegistry, candidate appointment.AppointmentUseCase) *AppointmentAPIHandler {
	return &AppointmentAPIHandler{
		router:       router,
		CoreRegistry: coreRegistry,
	}
}

func (h *AppointmentAPIHandler) Init() {
	router := h.router

	router.GET("/appointment/list", h.AppointmentList)
}

func (h *AppointmentAPIHandler) AppointmentList(c *gin.Context) {
	request := &models.AppointmentListRequest{}
	if err := c.Bind(request); err != nil {
		return
	}

	c.String(http.StatusOK, "User Info: %v", nil)
}
