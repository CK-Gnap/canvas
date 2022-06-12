package models

import (
	"time"

	"gorm.io/gorm"
)

type Canvas struct {
	gorm.Model
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func CreateCanvas(db *gorm.DB, Canvas *Canvas) (err error) {
	err = db.Create(Canvas).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCanvases(db *gorm.DB, Canvas *[]Canvas) (err error) {
	err = db.Find(Canvas).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCanvas(db *gorm.DB, Canvas *Canvas, id string) (err error) {
	err = db.Where("id = ?", id).First(Canvas).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateCanvas(db *gorm.DB, Canvas *Canvas) (err error) {
	db.Save(Canvas)
	return nil
}

func DeleteCanvas(db *gorm.DB, Canvas *Canvas, id string) (err error) {
	db.Where("id = ?", id).Delete(Canvas)
	return nil
}
