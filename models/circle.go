package models

import (
	models_interfaces "canvas/models/Interfaces"
	"math"
)

type Circle struct {
	Id       int64    `json:"id"`
	CanvasId int64    `json:"canvas_id"`
	Type     TypeEnum `json:"type"`
	X        float64  `json:"x"`
	Y        float64  `json:"y"`
	Radius   float64  `json:"radius"`
	Color    string   `json:"color"`
}

type CircleRequestCreate struct {
	Id       int64    `json:"id"`
	CanvasId int64    `json:"canvas_id"`
	Type     TypeEnum `json:"type"`
	X        float64  `json:"x"  binding:"required"`
	Y        float64  `json:"y"  binding:"required"`
	Radius   float64  `json:"radius" binding:"required"`
	Color    string   `json:"color"`
}

func ConvertToCircle(shape *Shape) models_interfaces.ShapeInterface {
	return &Circle{
		Id:       shape.Id,
		CanvasId: shape.CanvasId,
		Type:     shape.Type,
		X:        shape.X,
		Y:        shape.Y,
		Radius:   shape.Radius,
		Color:    shape.Color,
	}
}

func (Circle *Circle) GetType() string {
	return string(Circle.Type)
}

func (Circle *Circle) GetArea() float64 {
	return math.Pi * Circle.Radius * Circle.Radius
}

func (Circle *Circle) GetPerimeter() float64 {
	return 2 * math.Pi * Circle.Radius
}
