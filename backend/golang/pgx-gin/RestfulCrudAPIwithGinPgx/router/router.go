package router

import (
	"CrudAPIWithGin/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter(tagsController *controller.TagsControllerX) *gin.Engine {
	router := gin.Default()

	baseRouter := router.Group("/api")
	// tagsRouter := baseRouter.Group("/tags")
	// tagsRouter.GET("", tagsController.FindAll)
	// tagsRouter.GET("/:tagId", tagsController.FindById)
	// tagsRouter.POST("", tagsController.Create)
	// tagsRouter.PATCH("/:tagId", tagsController.Update)
	// tagsRouter.DELETE("/:tagId", tagsController.Delete)

	v1 := baseRouter.Group("/v1")
	// v1.GET("", tagsController.HelloX)
	v1.POST("", tagsController.CreateX)
	v1.GET("/:tagId", tagsController.FindByIdX)
	v1.GET("", tagsController.FindAllX)
	v1.PATCH("/:tagId", tagsController.UpdateX)
	v1.DELETE("/:tagId", tagsController.DeleteX)

	return router
}
