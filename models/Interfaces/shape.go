package models

type ShapeInterface interface {
	GetType() string
	GetArea() float64
	GetPerimeter() float64
}
