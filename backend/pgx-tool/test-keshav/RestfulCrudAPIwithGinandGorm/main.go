package main

import (
	"CrudAPIWithGin/config"
	"CrudAPIWithGin/controller"
	"CrudAPIWithGin/helper"
	"CrudAPIWithGin/model"
	"CrudAPIWithGin/repository"
	"CrudAPIWithGin/router"
	"CrudAPIWithGin/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Server Started")

	//database  k

	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	//Repository
	tagsRepository := repository.NewTagsRepositoryImpl(db)

	//Service
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)

	//Controller
	tagsController := controller.NewTagsController(tagsService)

	//router
	routes := router.NewRouter(tagsController)

	// routes := gin.Default()
	// routes.GET("", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, "Welcome")
	// })

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)

}
