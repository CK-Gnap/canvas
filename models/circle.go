package models

import (
	"math"
	"strconv"

	"gorm.io/gorm"
)

type Circle struct {
	Radius float64 `json:"radius"`
}

func (Circle *Circle) CreateShape(db *gorm.DB, Shape *Shape, canvasID string) (err error) {
	canvasId, _ := strconv.ParseInt(canvasID, 10, 64)
	Shape.CanvasId = canvasId
	Circle.Radius = Shape.Radius
	Shape.Area = Circle.GetArea()
	Shape.Perimeter = Circle.GetPerimeter()

	err = db.Create(Shape).Error
	if err != nil {
		return err
	}
	return nil
}

func (Circle *Circle) UpdateShape(db *gorm.DB, Shape *Shape) (err error) {
	Circle.Radius = Shape.Radius
	Shape.Area = Circle.GetArea()
	Shape.Perimeter = Circle.GetPerimeter()
	db.Save(Shape)
	return nil
}

func (Circle *Circle) GetArea() float64 {
	return math.Pi * Circle.Radius * Circle.Radius
}

func (Circle *Circle) GetPerimeter() float64 {
	return 2 * math.Pi * Circle.Radius
}
