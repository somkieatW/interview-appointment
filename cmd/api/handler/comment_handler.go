package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/somkieatW/interview-appointment/pkg/core/registry/core"
	"github.com/somkieatW/interview-appointment/pkg/core/utils"
	"github.com/somkieatW/interview-appointment/pkg/models"
	"github.com/somkieatW/interview-appointment/pkg/modules/comment"
)

type CommentAPIHandler struct {
	router         fiber.Router
	CoreRegistry   *core.CoreRegistry
	commentUseCase comment.CommentUseCase
}

func NewCommentAPIHandler(router fiber.Router, coreRegistry *core.CoreRegistry, commentUseCase comment.CommentUseCase) *CommentAPIHandler {
	return &CommentAPIHandler{
		router:         router,
		CoreRegistry:   coreRegistry,
		commentUseCase: commentUseCase,
	}
}

func (h *CommentAPIHandler) Init() {
	router := h.router

	router.Post("/comment/", h.Create)
}

func (h *CommentAPIHandler) Create(c *fiber.Ctx) error {
	request := &models.CommentCreateRequest{}
	if err := c.BodyParser(request); err != nil {
		if err != nil {
			response := utils.ErrorResponse{
				Message: "Internal server error",
			}
			log.Error(err)
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}
	}

	ctx := c.UserContext()
	data, err := h.commentUseCase.Create(ctx, request)
	if err != nil {
		response := utils.ErrorResponse{
			Message: "Internal server error",
		}
		log.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(data)

}
