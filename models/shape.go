package models

import (
	"time"

	"gorm.io/gorm"
)

type TypeEnum string

const (
	CIRCLE    TypeEnum = "circle"
	RECTANGLE TypeEnum = "rectangle"
	TRIANGLE  TypeEnum = "triangle"
)

type Shape struct {
	Id        int64     `json:"id"`
	X         float64   `json:"x"`
	Y         float64   `json:"y"`
	Width     float64   `json:"width"`
	Height    float64   `json:"height"`
	Radius    float64   `json:"radius"`
	SideLeft  float64   `json:"side_left"`
	SideRight float64   `json:"side_right"`
	SideBase  float64   `json:"side_base"`
	Color     string    `json:"color"`
	Area      float64   `json:"area"`
	Perimeter float64   `json:"perimeter"`
	Type      TypeEnum  `json:"type"`
	CanvasId  int64     `json:"canvas_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func GetShapes(db *gorm.DB, Shape *[]Shape, canvasID string) (err error) {
	err = db.Where("canvas_id  = ?", canvasID).Find(Shape).Error
	if err != nil {
		return err
	}
	return nil
}

func GetShape(db *gorm.DB, Shape *Shape, id string) (err error) {
	err = db.Where("id = ?", id).First(Shape).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteShape(db *gorm.DB, Shape *Shape, id string) (err error) {
	db.Where("id = ?", id).Delete(Shape)
	return nil
}
