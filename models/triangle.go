package models

import (
	models_interfaces "canvas/models/Interfaces"
)

type Triangle struct {
	Id       int64    `json:"id"`
	CanvasId int64    `json:"canvas_id"`
	Type     TypeEnum `json:"type"`
	X        float64  `json:"x"  binding:"required"`
	Y        float64  `json:"y"  binding:"required"`
	Width    float64  `json:"width"  binding:"required"`
	Height   float64  `json:"height"  binding:"required"`
	Color    string   `json:"color"`
}

type TriangleRequestCreate struct {
	Id       int64    `json:"id"`
	CanvasId int64    `json:"canvas_id"`
	Type     TypeEnum `json:"type"`
	X        float64  `json:"x"  binding:"required"`
	Y        float64  `json:"y"  binding:"required"`
	Width    float64  `json:"width"  binding:"required"`
	Height   float64  `json:"height"  binding:"required"`
	Color    string   `json:"color"`
}

func ConvertToTriangle(shape *Shape) models_interfaces.ShapeInterface {
	return &Triangle{
		Id:       shape.Id,
		CanvasId: shape.CanvasId,
		Type:     shape.Type,
		X:        shape.X,
		Y:        shape.Y,
		Width:    shape.Width,
		Height:   shape.Height,
		Color:    shape.Color,
	}
}

func (Triangle *Triangle) GetType() string {
	return string(Triangle.Type)
}

func (Triangle *Triangle) GetArea() float64 {
	return 0.5 * Triangle.Width * Triangle.Height
}

func (Triangle *Triangle) GetPerimeter() float64 {
	return Triangle.Width + (Triangle.Height * 2)
}
