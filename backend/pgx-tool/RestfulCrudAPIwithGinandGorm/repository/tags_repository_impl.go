package repository

import (
	"CrudAPIWithGin/data/request"
	"CrudAPIWithGin/helper"
	"CrudAPIWithGin/model"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"gorm.io/gorm"
)

// * pgx
type TagsRepositoryImplX struct {
	DbX *pgx.Conn
}

func NewTagsRepositoryImpX(conn *pgx.Conn) *TagsRepositoryImplX {
	return &TagsRepositoryImplX{DbX: conn}
}

type TagsRepositoryImpl struct {
	Db *gorm.DB
}

func NewTagsRepositoryImpl(Db *gorm.DB) TagRepository {
	return &TagsRepositoryImpl{Db: Db}
}

// Delete implements TagRepository.
func (t *TagsRepositoryImpl) Delete(tagsId int) {
	var tags model.Tags
	result := t.Db.Where("id = ?", tagsId).Delete(&tags)
	helper.ErrorPanic(result.Error)
}

// func (t *TagsRepositoryImpl) DeleteX(tagsId int) error {
// 	_, err := t.Db.Exec(context.Background(), "DELETE FROM tags WHERE id = $1", tagsId)
// 	return err
// }

// FindAll implements TagRepository.
func (t *TagsRepositoryImpl) FindAll() []model.Tags {
	var tags []model.Tags
	result := t.Db.Find(&tags)
	helper.ErrorPanic(result.Error)
	return tags
}

// FindById implements TagRepository.
func (t *TagsRepositoryImpl) FindById(tagsId int) (tags model.Tags, err error) {
	var tag model.Tags
	result := t.Db.Find(&tag, tagsId)
	if result != nil {
		return tag, nil
	} else {
		return tag, errors.New("tag is not found")
	}
}

// Save implements TagRepository.
func (t *TagsRepositoryImpl) Save(tags model.Tags) {
	result := t.Db.Create(&tags)
	helper.ErrorPanic(result.Error)
}

func (t *TagsRepositoryImplX) SaveX(tags model.TagsX) {
	_, err := t.DbX.Exec(context.Background(), "INSERT INTO tags (id, name) VALUES ($1, $2)", tags.Id, tags.Name)
	if err != nil {
		fmt.Print("dupid")
	}
	// return err
}

// Update implements TagRepository.
func (t *TagsRepositoryImpl) Update(tags model.Tags) {

	var updateTag = request.UpdateTagsRequest{
		Id:   tags.Id,
		Name: tags.Name,
	}

	result := t.Db.Model(&tags).Updates(updateTag)
	helper.ErrorPanic(result.Error)

}
