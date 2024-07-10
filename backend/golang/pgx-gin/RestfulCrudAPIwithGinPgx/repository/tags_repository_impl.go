package repository

import (
	"CrudAPIWithGin/model"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// * pgx
type TagsRepositoryImplX struct {
	DbX *pgx.Conn
}

func NewTagsRepositoryImpX(conn *pgx.Conn) *TagsRepositoryImplX {
	return &TagsRepositoryImplX{DbX: conn}
}

// Delete implements TagRepository.
func (t *TagsRepositoryImplX) DeleteX(tagsId int) {
	query := "DELETE FROM dupid WHERE id = $1"
	// Execute the query
	_, err := t.DbX.Exec(context.Background(), query, tagsId)
	if err != nil {
		fmt.Print("failed to delete tag")
	}
}

// FindById implements TagRepository.
func (t *TagsRepositoryImplX) FindByIdX(tagsId int) (tags model.TagsX, err error) {

	var tag model.TagsX
	err = t.DbX.QueryRow(context.Background(), "SELECT * FROM dupid WHERE id = $1", tagsId).Scan(&tag.Id, &tag.Name)
	if err != nil {
		if err == pgx.ErrNoRows {
			return tag, errors.New("tag is not found")
		}
		return tag, err
	}
	return tag, nil
}

// Save implements TagRepository.
func (t *TagsRepositoryImplX) SaveX(tags model.TagsX) {
	_, err := t.DbX.Exec(context.Background(), "INSERT INTO dupid (name) VALUES ($1)", tags.Name)
	if err != nil {
		fmt.Print("dupid")
	}
}

// Update implements TagRepository.
func (t *TagsRepositoryImplX) UpdateX(tags model.TagsX) {
	// Define the SQL update query
	query := "UPDATE dupid SET name = $1  WHERE id = $2"

	// Execute the query
	_, err := t.DbX.Exec(context.Background(), query, tags.Name, tags.Id)
	if err != nil {
		fmt.Print("dupid")
	}
}

// FindAll implements TagRepository.
func (t *TagsRepositoryImplX) FindAllX() []model.TagsX {
	// Define the SQL select query
	query := "SELECT * FROM dupid"

	// Execute the query
	rows, err := t.DbX.Query(context.Background(), query)
	if err != nil {
		fmt.Print("failed to find all tags")
	}
	defer rows.Close()

	var tags []model.TagsX
	for rows.Next() {
		var tag model.TagsX
		err := rows.Scan(&tag.Id, &tag.Name)
		if err != nil {

			fmt.Print("failed to scan tag")
		}
		tags = append(tags, tag)
	}

	if rows.Err() != nil {

		fmt.Print("error during rows iteration")
	}

	return tags
}

// type TagsRepositoryImpl struct {
// 	Db *gorm.DB
// }

// func NewTagsRepositoryImpl(Db *gorm.DB) TagRepository {
// 	return &TagsRepositoryImpl{Db: Db}
// }

// func (t *TagsRepositoryImpl) Delete(tagsId int) {
// 	var tags model.Tags
// 	result := t.Db.Where("id = ?", tagsId).Delete(&tags)
// 	helper.ErrorPanic(result.Error)
// }

// func (t *TagsRepositoryImpl) FindAll() []model.Tags {
// 	var tags []model.Tags
// 	result := t.Db.Find(&tags)
// 	helper.ErrorPanic(result.Error)
// 	return tags
// }

// func (t *TagsRepositoryImpl) FindById(tagsId int) (tags model.Tags, err error) {
// 	var tag model.Tags
// 	result := t.Db.Find(&tag, tagsId)
// 	if result != nil {
// 		return tag, nil
// 		} else {
// 			return tag, errors.New("tag is not found")
// 		}
// 	}

// func (t *TagsRepositoryImpl) Save(tags model.Tags) {
// 	result := t.Db.Create(&tags)
// 	helper.ErrorPanic(result.Error)
// }

// func (t *TagsRepositoryImpl) Update(tags model.Tags) {

// 	var updateTag = request.UpdateTagsRequest{
// 		Id:   tags.Id,
// 		Name: tags.Name,
// 	}

// 	result := t.Db.Model(&tags).Updates(updateTag)
// 	helper.ErrorPanic(result.Error)
// }
