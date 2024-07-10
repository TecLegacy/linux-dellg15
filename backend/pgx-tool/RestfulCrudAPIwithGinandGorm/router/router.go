package router

import (
	"CrudAPIWithGin/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(tagsController *controller.TagsControllerX) *gin.Engine {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")
	// tagsRouter := baseRouter.Group("/tags")
	// tagsRouter.GET("", tagsController.FindAll)
	// tagsRouter.GET("/:tagId", tagsController.FindById)
	// tagsRouter.POST("", tagsController.Create)
	// tagsRouter.PATCH("/:tagId", tagsController.Update)
	// tagsRouter.DELETE("/:tagId", tagsController.Delete)

	v1 := baseRouter.Group("/v1")
	v1.GET("", tagsController.HelloX)
	v1.POST("", tagsController.CreateX)

	return router
}
