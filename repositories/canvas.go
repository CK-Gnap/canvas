package repositories

import (
	"canvas/database"
	"canvas/models"
	repositories_interfaces "canvas/repositories/Interfaces"

	"gorm.io/gorm"
)

type CanvasRepo struct {
	db *gorm.DB
}

func NewCanvasRepo(db *gorm.DB) repositories_interfaces.CanvasRepoInterface {
	return &CanvasRepo{db}
}

func (repo *CanvasRepo) CreateCanvas(Canvas *models.Canvas) (err error) {
	err = database.Db.Create(Canvas).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *CanvasRepo) GetCanvases(Canvas *[]models.Canvas) (err error) {
	err = database.Db.Find(Canvas).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *CanvasRepo) GetCanvas(Canvas *models.Canvas, id string) (err error) {
	err = database.Db.Where("id = ?", id).First(Canvas).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *CanvasRepo) UpdateCanvas(Canvas *models.Canvas, id string) (err error) {
	database.Db.Save(Canvas)
	return nil
}

func (repo *CanvasRepo) DeleteCanvas(Canvas *models.Canvas, id string) (err error) {
	database.Db.Where("id = ?", id).Delete(Canvas)
	return nil
}
