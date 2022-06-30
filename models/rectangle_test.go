package models

import (
	models_interfaces "canvas/models/Interfaces"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertToRectangle(t *testing.T) {
	tests := []struct {
		name  string
		shape *Shape
		want  models_interfaces.ShapeInterface
	}{
		{
			name: "when happy",
			shape: &Shape{
				Id:       1,
				CanvasId: 1,
				Type:     RECTANGLE,
				X:        1,
				Y:        1,
				Width:    1,
				Height:   1,
				Color:    "#000000",
			},
			want: &Rectangle{
				Id:       1,
				CanvasId: 1,
				Type:     RECTANGLE,
				X:        1,
				Y:        1,
				Width:    1,
				Height:   1,
				Color:    "#000000",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ConvertToRectangle(test.shape)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestGetTypeRectangle(t *testing.T) {
	tests := []struct {
		name      string
		shapeType TypeEnum
		want      string
	}{
		{
			name:      "when happy",
			shapeType: RECTANGLE,
			want:      string(RECTANGLE),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			shape := Rectangle{
				Type: test.shapeType,
			}
			got := shape.GetType()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestGetAreaRectangle(t *testing.T) {
	tests := []struct {
		name   string
		Width  float64
		Height float64
		want   float64
	}{
		{
			name:   "when happy",
			Width:  10,
			Height: 10,
			want:   100,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			shape := Rectangle{
				Width:  test.Width,
				Height: test.Height,
			}
			got := shape.GetArea()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestGetPerimeterRectangle(t *testing.T) {
	tests := []struct {
		name   string
		Width  float64
		Height float64
		want   float64
	}{
		{
			name:   "when happy",
			Width:  10,
			Height: 10,
			want:   200,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			shape := Rectangle{
				Width:  test.Width,
				Height: test.Height,
			}
			got := shape.GetPerimeter()
			assert.Equal(t, test.want, got)
		})
	}
}
