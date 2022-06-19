package models

import (
	"errors"
	"math"
)

type Triangle struct {
	CanvasId  int64   `json:"canvas_id"`
	X         float64 `json:"x"  binding:"required"`
	Y         float64 `json:"y"  binding:"required"`
	Width     float64 `json:"width"  binding:"required"`
	Height    float64 `json:"height"  binding:"required"`
	SideLeft  float64 `json:"sideLeft"`
	SideRight float64 `json:"sideRight"`
	SideBase  float64 `json:"sideBase"`
	Color     string  `json:"color"`
	Area      float64 `json:"area"`
	Perimeter float64 `json:"perimeter"`
}

func (Triangle *Triangle) GetArea() float64 {
	return 0.5 * Triangle.SideBase * Triangle.Height
}

func (Triangle *Triangle) GetPerimeter() float64 {
	return Triangle.SideLeft + Triangle.SideRight + Triangle.SideBase
}

func (Triangle *Triangle) GetSides() {
	width := Triangle.Width
	height := Triangle.Height
	sideA := math.Pow(height, 2)
	sideB := math.Pow(width, 2)
	sideC := math.Sqrt(sideA + sideB)

	if width > height {
		Triangle.SideLeft = float64(height)
		Triangle.SideRight = float64(sideC)
		Triangle.SideBase = float64(width)
	} else {
		Triangle.SideLeft = float64(height)
		Triangle.SideRight = float64(width)
		Triangle.SideBase = float64(sideC)
	}
}

func (Triangle *Triangle) CheckIsTriangle() (err error) {
	sideLeft := Triangle.SideLeft
	sideRight := Triangle.SideRight
	sideBase := Triangle.SideBase
	if (sideLeft+sideRight > sideBase) && (sideLeft+sideBase > sideRight) && (sideRight+sideBase > sideLeft) {
		return nil
	}
	return errors.New("Not a triangle")
}
