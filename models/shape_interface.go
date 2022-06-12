package models

import "gorm.io/gorm"

type ShapeInterface interface {
	CreateShape(db *gorm.DB, Shape *Shape, canvasID string) (err error)
	UpdateShape(db *gorm.DB, Shape *Shape) (err error)
	GetArea() float64
	GetPerimeter() float64
}
