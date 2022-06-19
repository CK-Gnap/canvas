package models

type Rectangle struct {
	CanvasId  int64   `json:"canvas_id"`
	X         float64 `json:"x"  binding:"required"`
	Y         float64 `json:"y"  binding:"required"`
	Width     float64 `json:"width"  binding:"required"`
	Height    float64 `json:"height"  binding:"required"`
	Color     string  `json:"color"`
	Area      float64 `json:"area"`
	Perimeter float64 `json:"perimeter"`
}

func (Rectangle *Rectangle) GetArea() float64 {
	return Rectangle.Width * Rectangle.Height
}

func (Rectangle *Rectangle) GetPerimeter() float64 {
	return 2 * (Rectangle.Width * Rectangle.Height)
}
