package main

import (
	"CrudAPIWithGin/config"
	"CrudAPIWithGin/controller"
	"CrudAPIWithGin/helper"
	"os"

	// "CrudAPIWithGin/model"
	"CrudAPIWithGin/repository"
	"CrudAPIWithGin/router"
	"CrudAPIWithGin/service"

	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Server Started")

	// db := config.DatabaseConnection()

	//* pgx
	connX := config.DatabaseConnectionX()
	defer connX.Close(context.Background())

	// Create tags table
	if err := config.CreateTagsTable(connX); err != nil {
		log.Fatal().Err(err).Msg("Failed to create tags table")
		os.Exit(1)
	}

	validate := validator.New()

	// db.Table("tags").AutoMigrate(&model.Tags{})

	//Repository

	// tagsRepository := repository.NewTagsRepositoryImpl(db)

	tagsRepositoryX := repository.NewTagsRepositoryImpX(connX)

	//Service
	// tagsService := service.NewTagsServiceImpl(tagsRepository, validate)

	tagsServiceX := service.NewTagsServiceImplX(tagsRepositoryX, validate)

	//Controller
	tagsController := controller.NewTagsControllerX(tagsServiceX)

	//router
	routes := router.NewRouter(tagsController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)

}
