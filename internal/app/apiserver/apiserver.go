package apiserver

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	contentrestapi "github.com/something-to-start-with/api-server-go/internal/content/controller/restapi"
	contentrepository "github.com/something-to-start-with/api-server-go/internal/content/repository/postgres"
	contentservice "github.com/something-to-start-with/api-server-go/internal/content/service"
	"github.com/something-to-start-with/api-server-go/internal/pkg/db/postgres"
)

func Run(cfg *Config) {

	db, err := postgres.New(cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Error closing connect to database: %s", err)
		}
	}(db)

	router := gin.Default()
	setupContentRoutes(router, db)

	log.Fatal(router.Run(":" + cfg.Server.Port))
}

func setupContentRoutes(router *gin.Engine, db *sqlx.DB) {
	contentRepository := contentrepository.New(db)
	contentService := contentservice.New(contentRepository)
	contentrestapi.SetupRoutes(router, contentService)
}
