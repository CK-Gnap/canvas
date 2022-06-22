package models

import (
	models "canvas/models/Interfaces"
	"time"
)

type Canvas struct {
	Id        int64                   `json:"id"`
	Name      string                  `json:"name"`
	Width     float64                 `json:"width"`
	Height    float64                 `json:"height"`
	Color     string                  `json:"color"`
	Shapes    []models.ShapeInterface `json:"shapes" gorm:"-"`
	CreatedAt time.Time               `json:"created_at"`
	UpdatedAt time.Time               `json:"updated_at"`
}

func (Canvas *Canvas) TableName() string {
	return "canvas"
}

type CanvasRequestCreate struct {
	Name   string  `json:"name"  binding:"required"`
	Width  float64 `json:"width"  binding:"required"`
	Height float64 `json:"height"  binding:"required"`
	Color  string  `json:"color" binding:"required"`
}

type CanvasRequestUpdare struct {
	Name   string  `json:"name"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Color  string  `json:"color"`
}
