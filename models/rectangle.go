package models

import (
	models_interfaces "canvas/models/Interfaces"
)

type Rectangle struct {
	Id       int64    `json:"id"`
	CanvasId int64    `json:"canvas_id"`
	Type     TypeEnum `json:"type"`
	X        float64  `json:"x"`
	Y        float64  `json:"y"`
	Width    float64  `json:"width"`
	Height   float64  `json:"height"`
	Color    string   `json:"color"`
}

type RectangleRequestCreate struct {
	Id       int64    `json:"id"`
	CanvasId int64    `json:"canvas_id"`
	Type     TypeEnum `json:"type"`
	X        float64  `json:"x"  binding:"required"`
	Y        float64  `json:"y"  binding:"required"`
	Width    float64  `json:"width"  binding:"required"`
	Height   float64  `json:"height"  binding:"required"`
	Color    string   `json:"color"`
}

func ConvertToRectangle(shape *Shape) models_interfaces.ShapeInterface {
	return &Rectangle{
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

func (Rectangle *Rectangle) GetType() string {
	return string(Rectangle.Type)
}

func (Rectangle *Rectangle) GetArea() float64 {
	return Rectangle.Width * Rectangle.Height
}

func (Rectangle *Rectangle) GetPerimeter() float64 {
	return 2 * (Rectangle.Width * Rectangle.Height)
}
