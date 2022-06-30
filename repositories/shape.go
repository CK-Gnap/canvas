package repositories

import (
	"canvas/constants"
	"canvas/models"
	repositories_interfaces "canvas/repositories/Interfaces"

	"gorm.io/gorm"
)

type ShapeRepo struct {
	db *gorm.DB
}

func NewShapeRepo(db *gorm.DB) repositories_interfaces.ShapeRepoInterface {
	return &ShapeRepo{db}
}

func (repo *ShapeRepo) CreateShape(Shape *models.Shape, canvasID string) (err error) {
	err = repo.db.Create(Shape).Error
	if err != nil {
		return constants.ErrCreateShape
	}
	return nil
}

func (repo *ShapeRepo) GetShapes(Shape *[]models.Shape, canvasID string) (*[]models.Shape, error) {
	err := repo.db.Where("canvas_id  = ?", canvasID).Find(Shape).Error
	if err != nil {
		return nil, constants.ErrGetShapes
	}
	return Shape, nil
}

func (repo *ShapeRepo) GetShape(Shape *models.Shape, id string) (err error) {
	err = repo.db.Where("id = ?", id).First(Shape).Error
	if err != nil {
		return constants.ErrGetShape
	}
	return nil
}

func (repo *ShapeRepo) UpdateShape(Shape *models.Shape, id string) (*models.Shape, error) {
	err := repo.db.Save(Shape).Error
	if err != nil {
		return nil, constants.ErrUpdateShape
	}
	return Shape, nil
}

func (repo *ShapeRepo) DeleteShape(Shape *models.Shape, id string) (err error) {
	err = repo.db.Where("id = ?", id).Delete(Shape).Error
	if err != nil {
		return constants.ErrDeleteShape
	}
	return nil
}
