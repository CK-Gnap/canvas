package models

import (
	"canvas/database"
	"time"
)

type Canvas struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"  binding:"required"`
	Shape     *[]Shape  `json:"shape"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateCanvas(Canvas *Canvas) (err error) {
	err = database.Db.Create(Canvas).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCanvas(Canvas *[]Canvas) (err error) {
	err = database.Db.Preload("Shape").Find(Canvas).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCanvasById(Canvas *Canvas, id string) (err error) {
	err = database.Db.Where("id = ?", id).Preload("Shape").First(Canvas).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateCanvas(Canvas *Canvas) (err error) {
	database.Db.Save(Canvas)
	return nil
}

func DeleteCanvas(Canvas *Canvas, id string) (err error) {
	database.Db.Where("id = ?", id).Delete(Canvas)
	return nil
}
