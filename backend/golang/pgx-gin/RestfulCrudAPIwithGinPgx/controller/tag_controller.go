package controller

import (
	"CrudAPIWithGin/data/request"
	"CrudAPIWithGin/data/response"
	"CrudAPIWithGin/helper"
	"CrudAPIWithGin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TagsController struct {
	tagsService service.TagsService
}

func NewTagsController(service service.TagsService) *TagsController {
	return &TagsController{
		tagsService: service,
	}
}

// *PGX
type TagsControllerX struct {
	tagsServiceX service.TagsServiceX
}

func NewTagsControllerX(service service.TagsServiceX) *TagsControllerX {
	return &TagsControllerX{
		tagsServiceX: service,
	}
}

func (controller *TagsControllerX) HelloX(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   nil,
	})
}

func (controller *TagsControllerX) CreateX(ctx *gin.Context) {
	createTagsRequest := request.CreateTagsRequest{}
	err := ctx.ShouldBindJSON(&createTagsRequest)
	helper.ErrorPanic(err)

	//* Create tags with PGX
	controller.tagsServiceX.CreateX(createTagsRequest)

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   nil,
	})
}

func (controller *TagsControllerX) FindByIdX(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	tagResponse := controller.tagsServiceX.FindByIdX(id)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *TagsControllerX) UpdateX(ctx *gin.Context) {
	updateTagsRequest := request.UpdateTagsRequest{}
	err := ctx.ShouldBindJSON(&updateTagsRequest)
	helper.ErrorPanic(err)
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	updateTagsRequest.Id = id

	controller.tagsServiceX.UpdateX(updateTagsRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
func (controller *TagsControllerX) DeleteX(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	controller.tagsServiceX.DeleteX(id)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TagsControllerX) FindAllX(ctx *gin.Context) {
	tagResponse := controller.tagsServiceX.FindAllX()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

/*
func (controller *TagsController) FindById(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	tagResponse := controller.tagsService.FindById(id)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

*/