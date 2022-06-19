package models

import (
	"math"
)

type Circle struct {
	CanvasId  int64   `json:"canvas_id"`
	X         float64 `json:"x"  binding:"required"`
	Y         float64 `json:"y"  binding:"required"`
	Radius    float64 `json:"radius" binding:"required"`
	Color     string  `json:"color"`
	Area      float64 `json:"area"`
	Perimeter float64 `json:"perimeter"`
}

func (Circle *Circle) GetArea() float64 {
	return math.Pi * Circle.Radius * Circle.Radius
}

func (Circle *Circle) GetPerimeter() float64 {
	return 2 * math.Pi * Circle.Radius
}
