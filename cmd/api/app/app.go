package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/somkieatW/interview-appointment/cmd/api/handler"
	"github.com/somkieatW/interview-appointment/pkg/config"
	"github.com/somkieatW/interview-appointment/pkg/core/db"
	"github.com/somkieatW/interview-appointment/pkg/core/registry"
	"github.com/somkieatW/interview-appointment/pkg/core/registry/core"
	candidate "github.com/somkieatW/interview-appointment/pkg/modules/appointment/usecase"
	comment "github.com/somkieatW/interview-appointment/pkg/modules/comment/usecase"
	"github.com/somkieatW/interview-appointment/pkg/repository"
	"log"
)

func Run() {
	config.Init("pkg/config")
	cfg := config.LoadConfig()

	mysqlDb, err := db.NewMySqlDB(cfg.Mysql, cfg.Secret)
	if err != nil {
		log.Panicf("new mysql db: %v", err)
	}
	defer func() {
		if err := mysqlDb.Close(); err != nil {
			log.Panicf("close mysql db: %v", err)
		}
	}()

	repositoryRegistry := &registry.RepositoryRegistry{
		AppointmentRepository: repository.NewAppointmentRepository(mysqlDb.Client),
		UserRepository:        repository.NewUserRepository(mysqlDb.Client),
		CommentRepository:     repository.NewCommentRepository(mysqlDb.Client),
	}

	coreRegistry := &core.CoreRegistry{
		Config: cfg,
	}

	candidateUseCase := candidate.NewAppointmentUseCase(coreRegistry, repositoryRegistry)
	commentUseCase := comment.NewCommentUseCase(coreRegistry, repositoryRegistry)

	r := fiber.New()
	router := r.Group("")

	handler.NewAppointmentAPIHandler(router, coreRegistry, candidateUseCase).Init()
	handler.NewCommentAPIHandler(router, coreRegistry, commentUseCase).Init()

	r.Get("/health-check", func(c *fiber.Ctx) error {
		return c.JSON(c.App().Stack())
	})

	port := "3001"
	err = r.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Panic(err)
		return
	}
	var ()
}
