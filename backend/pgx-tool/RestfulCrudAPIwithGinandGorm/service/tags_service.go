package service

import (
	"CrudAPIWithGin/data/request"
	"CrudAPIWithGin/data/response"
)

type TagsService interface {
	Create(tags request.CreateTagsRequest)
	Update(tags request.UpdateTagsRequest)
	Delete(tagsId int)
	FindById(tagsId int) response.TagsResponse
	FindAll() []response.TagsResponse
}

type TagsServiceX interface {
	CreateX(tags request.CreateTagsRequest)
	// UpdateX(tags request.UpdateTagsRequest)
	// DeleteX(tagsId int)
	// FindByIdX(tagsId int) response.TagsResponse
	// FindAllX() []response.TagsResponse
}
