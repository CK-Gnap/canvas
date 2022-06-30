package models

import (
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
	Color     string    `json:"color"`
	Type      TypeEnum  `json:"type"`
	CanvasId  int64     `json:"canvas_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Shape *Shape) TableName() string {
	return "shape"
}

type ShapeRequestUpdate struct {
	X      float64 `json:"x"  binding:"required"`
	Y      float64 `json:"y"  binding:"required"`
	Width  float64 `json:"width" `
	Height float64 `json:"height"`
	Radius float64 `json:"radius"`
	Color  string  `json:"color"`
}
