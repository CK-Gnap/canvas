package repositories

import (
	"canvas/models"
	repositories_interfaces "canvas/repositories/Interfaces"
	"errors"

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
		return errors.New("Error creating shape")
	}
	return nil
}

func (repo *ShapeRepo) GetShapes(Shape *[]models.Shape, canvasID string) (*[]models.Shape, error) {
	err := repo.db.Where("canvas_id  = ?", canvasID).Find(Shape).Error
	if err != nil {
		return nil, errors.New("Error getting shapes")
	}
	return Shape, nil
}

func (repo *ShapeRepo) GetShape(Shape *models.Shape, id string) (err error) {
	err = repo.db.Where("id = ?", id).First(Shape).Error
	if err != nil {
		return errors.New("Error getting shape")
	}
	return nil
}

func (repo *ShapeRepo) UpdateShape(Shape *models.Shape, id string) (*models.Shape, error) {
	err := repo.db.Save(Shape).Error
	if err != nil {
		return nil, errors.New("Error updating shape")
	}
	return Shape, nil
}

func (repo *ShapeRepo) DeleteShape(Shape *models.Shape, id string) (err error) {
	err = repo.db.Where("id = ?", id).Delete(Shape).Error
	if err != nil {
		return errors.New("Error deleting shape")
	}
	return nil
}
