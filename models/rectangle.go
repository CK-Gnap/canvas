package models

import (
	"strconv"

	"gorm.io/gorm"
)

type Rectangle struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

func (Rectangle *Rectangle) CreateShape(db *gorm.DB, Shape *Shape, canvasID string) (err error) {
	canvasId, _ := strconv.ParseInt(canvasID, 10, 64)
	Shape.CanvasId = canvasId
	Rectangle.Width = Shape.Width
	Rectangle.Height = Shape.Height
	Shape.Area = Rectangle.GetArea()
	Shape.Perimeter = Rectangle.GetPerimeter()

	err = db.Create(Shape).Error
	if err != nil {
		return err
	}
	return nil
}

func (Rectangle *Rectangle) UpdateShape(db *gorm.DB, Shape *Shape) (err error) {
	Rectangle.Width = Shape.Width
	Rectangle.Height = Shape.Height
	Shape.Area = Rectangle.GetArea()
	Shape.Perimeter = Rectangle.GetPerimeter()
	db.Save(Shape)
	return nil
}

func (Rectangle *Rectangle) GetArea() float64 {
	return Rectangle.Width * Rectangle.Height
}

func (Rectangle *Rectangle) GetPerimeter() float64 {
	return 2 * (Rectangle.Width * Rectangle.Height)
}
