package models

import (
	"canvas/database"
	"time"
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
	Canvas    Canvas    `json:"canvas"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateShape(Shape *Shape, canvasID string) (err error) {
	err = database.Db.Create(Shape).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateShape(Shape *Shape) (err error) {
	database.Db.Save(Shape)
	return nil
}

func GetShapes(Shape *[]Shape, canvasID string) (err error) {
	err = database.Db.Where("canvas_id  = ?", canvasID).Preload("Canvas").Find(Shape).Error
	if err != nil {
		return err
	}
	return nil
}

func GetShape(Shape *Shape, id string) (CanvasId int64, Type string, err error) {
	err = database.Db.Where("id = ?", id).Preload("Canvas").First(Shape).Error
	if err != nil {
		return 0, "", err
	}
	return Shape.CanvasId, string(Shape.Type), nil
}

func DeleteShape(Shape *Shape, id string) (err error) {
	database.Db.Where("id = ?", id).Delete(Shape)
	return nil
}
