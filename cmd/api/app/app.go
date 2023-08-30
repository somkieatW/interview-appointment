package app

import (
	"github.com/gin-gonic/gin"
	"github.com/somkieatW/candidate/cmd/api/handler"
	"github.com/somkieatW/candidate/pkg/config"
	"github.com/somkieatW/candidate/pkg/core/db"
	"github.com/somkieatW/candidate/pkg/core/registry"
	"github.com/somkieatW/candidate/pkg/core/registry/core"
	candidate "github.com/somkieatW/candidate/pkg/modules/appointment/usecase"
	"github.com/somkieatW/candidate/pkg/repository"
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
	}

	coreRegistry := &core.CoreRegistry{
		Config: cfg,
	}

	candidateUseCase := candidate.NewAppointmentUseCase(coreRegistry, repositoryRegistry)

	r := gin.Default()
	router := r.Group("")

	handler.NewAppointmentAPIHandler(router, coreRegistry, candidateUseCase).Init()

	err = r.Run(":8080")
	if err != nil {
		log.Panicf("run api server err : %v", err)
	}

	var ()
}
