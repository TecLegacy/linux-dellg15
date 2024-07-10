package repository

import "CrudAPIWithGin/model"

type TagRepository interface {
	Save(tags model.Tags)
	Update(tags model.Tags)
	Delete(tagsId int)
	FindById(tagsId int) (tags model.Tags, err error)
	FindAll() []model.Tags
}

type TagRepositoryX interface {
	SaveX(tags model.TagsX)
	// UpdateX(tags model.Tags)
	// DeleteX(tagsId int)
	// FindByIdX(tagsId int) (tags model.Tags, err error)
	// FindAllX() []model.Tags
}
