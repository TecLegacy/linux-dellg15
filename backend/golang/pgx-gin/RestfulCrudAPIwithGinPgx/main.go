package main

import (
	"CrudAPIWithGin/config"
	"CrudAPIWithGin/controller"
	"CrudAPIWithGin/helper"
	"os"

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

	//* pgx
	connX := config.DatabaseConnectionX()
	defer connX.Close(context.Background())

	validate := validator.New()

	// Create tags table
	if err := config.CreateTagsTable(connX); err != nil {
		log.Fatal().Err(err).Msg("Failed to create tags table")
		os.Exit(1)
	}

	//Repository
	tagsRepositoryX := repository.NewTagsRepositoryImpX(connX)

	//Service
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
