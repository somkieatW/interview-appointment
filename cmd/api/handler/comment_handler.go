package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/somkieatW/candidate/pkg/core/registry/core"
	"github.com/somkieatW/candidate/pkg/modules/comment"
	"github.com/somkieatW/candidate/pkg/modules/comment/models"
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
		return err
	}

	ctx := c.UserContext()
	data, err := h.commentUseCase.Create(ctx, request)
	if err != nil {
		return err
	}

	c.Status(200).JSON(data)
	return nil
}
