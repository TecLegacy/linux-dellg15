package service

import (
	"CrudAPIWithGin/data/request"
	"CrudAPIWithGin/data/response"
	"CrudAPIWithGin/helper"
	"CrudAPIWithGin/model"
	"CrudAPIWithGin/repository"

	"github.com/go-playground/validator/v10"
)

// !PGX
type TagsServiceImplX struct {
	TagsRepositoryX repository.TagRepositoryX
	validate        *validator.Validate
}

func NewTagsServiceImplX(tagRepositoryX repository.TagRepositoryX, validate *validator.Validate) TagsServiceX {
	return &TagsServiceImplX{
		TagsRepositoryX: tagRepositoryX,
		validate:        validate,
	}
}

// Create implements TagsService.
func (t *TagsServiceImplX) CreateX(tags request.CreateTagsRequest) {
	err := t.validate.Struct(tags)
	helper.ErrorPanic(err)
	tagModel := model.TagsX{
		Name: tags.Name,
	}
	t.TagsRepositoryX.SaveX(tagModel)
}

// Delete implements TagsService.
func (t *TagsServiceImplX) DeleteX(tagsId int) {
	t.TagsRepositoryX.DeleteX(tagsId)
}

// FindAll implements TagsService.
func (t *TagsServiceImplX) FindAllX() []response.TagsResponse {
	result := t.TagsRepositoryX.FindAllX()
	var tags []response.TagsResponse
	for _, value := range result {
		tag := response.TagsResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		tags = append(tags, tag)
	}
	return tags
}

// FindById implements TagsService.
func (t *TagsServiceImplX) FindByIdX(tagsId int) response.TagsResponse {
	tagData, err := t.TagsRepositoryX.FindByIdX(tagsId)
	helper.ErrorPanic(err)
	tagResponse := response.TagsResponse{
		Id:   tagData.Id,
		Name: tagData.Name,
	}
	return tagResponse
}

// Update implements TagsService.
func (t *TagsServiceImplX) UpdateX(tags request.UpdateTagsRequest) {
	tagData, err := t.TagsRepositoryX.FindByIdX(tags.Id)
	helper.ErrorPanic(err)
	tagData.Name = tags.Name
	t.TagsRepositoryX.UpdateX(tagData)
}

// type TagsServiceImpl struct {
// 	TagsRepository repository.TagRepository
// 	validate       *validator.Validate
// }

// func NewTagsServiceImpl(tagRepository repository.TagRepository, validate *validator.Validate) TagsService {
// 	return &TagsServiceImpl{
// 		TagsRepository: tagRepository,
// 		validate:       validate,
// 	}
// }

// Create implements TagsService.
// func (t *TagsServiceImpl) Create(tags request.CreateTagsRequest) {
// 	err := t.validate.Struct(tags)
// 	helper.ErrorPanic(err)
// 	tagModel := model.Tags{
// 		Name: tags.Name,
// 	}
// 	t.TagsRepository.Save(tagModel)
// }

// func (t *TagsServiceImpl) Delete(tagsId int) {
// 	t.TagsRepository.Delete(tagsId)
// }

// // FindAll implements TagsService.
//
//	func (t *TagsServiceImpl) FindAll() []response.TagsResponse {
//		result := t.TagsRepository.FindAll()
//		var tags []response.TagsResponse
//		for _, value := range result {
//			tag := response.TagsResponse{
//				Id:   value.Id,
//				Name: value.Name,
//			}
//			tags = append(tags, tag)
//		}
//		return tags
//	}
//

//	func (t *TagsServiceImpl) FindById(tagsId int) response.TagsResponse {
//		tagData, err := t.TagsRepository.FindById(tagsId)
//		helper.ErrorPanic(err)
//		tagResponse := response.TagsResponse{
//			Id:   tagData.Id,
//			Name: tagData.Name,
//		}
//		return tagResponse
//	}

// func (t *TagsServiceImpl) Update(tags request.UpdateTagsRequest) {
// 	tagData, err := t.TagsRepository.FindById(tags.Id)
// 	helper.ErrorPanic(err)
// 	tagData.Name = tags.Name
// 	t.TagsRepository.Update(tagData)
// }
