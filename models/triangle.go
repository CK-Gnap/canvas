package models

import (
	"errors"
	"strconv"

	"gorm.io/gorm"
)

type Triangle struct {
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Width     float64 `json:"width"`
	Height    float64 `json:"height"`
	SideLeft  float64 `json:"sideLeft"`
	SideRight float64 `json:"sideRight"`
	SideBase  float64 `json:"sideBase"`
}

func (Triangle *Triangle) CreateShape(db *gorm.DB, Shape *Shape, canvasID string) (err error) {

	canvasId, _ := strconv.ParseInt(canvasID, 10, 64)
	Shape.CanvasId = canvasId
	Triangle.X = Shape.X
	Triangle.Y = Shape.Y
	Triangle.Width = Shape.Width
	Triangle.Height = Shape.Height

	Triangle.getSides()
	Shape.SideLeft = Triangle.SideLeft
	Shape.SideRight = Triangle.SideRight
	Shape.SideBase = Triangle.SideBase
	Shape.Area = Triangle.GetArea()
	Shape.Perimeter = Triangle.GetPerimeter()

	err = Triangle.checkIsTriangle()
	if err != nil {
		return err
	}

	err = db.Create(Shape).Error
	if err != nil {
		return err
	}
	return nil
}

func (Triangle *Triangle) UpdateShape(db *gorm.DB, Shape *Shape) (err error) {
	Triangle.Height = Shape.Height
	Triangle.X = Shape.X
	Triangle.Y = Shape.Y
	Triangle.Width = Shape.Width
	Triangle.Height = Shape.Height
	Triangle.getSides()
	Shape.SideLeft = Triangle.SideLeft
	Shape.SideRight = Triangle.SideRight
	Shape.SideBase = Triangle.SideBase
	Shape.Area = Triangle.GetArea()
	Shape.Perimeter = Triangle.GetPerimeter()

	err = Triangle.checkIsTriangle()
	if err != nil {
		return err
	}

	db.Save(Shape)
	return nil
}

func (Triangle *Triangle) GetArea() float64 {
	return 0.5 * Triangle.SideBase * Triangle.Height
}

func (Triangle *Triangle) GetPerimeter() float64 {
	return Triangle.SideLeft + Triangle.SideRight + Triangle.SideBase
}

func (Triangle *Triangle) getSides() {
	x := int(Triangle.X)
	y := int(Triangle.Y)
	width := int(Triangle.Width)
	height := int(Triangle.Height)
	line := 0
	baseRadius := 0
	sideBase := 0
	sideLeft := 0
	sideRight := 0

	for i := y; i <= y+height; i++ {
		line++
		sideLeft = width + line
		sideRight = width - line
		baseRadius = i
	}
	for ii := x; ii <= baseRadius; ii++ {
		sideBase = baseRadius * 2
	}
	Triangle.SideLeft = float64(sideLeft)
	Triangle.SideRight = float64(sideRight)
	Triangle.SideBase = float64(sideBase)
}

func (Triangle *Triangle) checkIsTriangle() (err error) {
	sideLeft := Triangle.SideLeft
	sideRight := Triangle.SideRight
	sideBase := Triangle.SideBase
	if (sideLeft+sideRight > sideBase) && (sideLeft+sideBase > sideRight) && (sideRight+sideBase > sideLeft) {
		return nil
	}
	return errors.New("Not a triangle")
}
