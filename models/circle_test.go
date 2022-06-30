package models

import (
	models_interfaces "canvas/models/Interfaces"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertToCircle(t *testing.T) {
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
				Type:     CIRCLE,
				X:        10,
				Y:        10,
				Radius:   10,
				Color:    "#ffffff",
			},
			want: &Circle{
				Id:       1,
				CanvasId: 1,
				Type:     CIRCLE,
				X:        10,
				Y:        10,
				Radius:   10,
				Color:    "#ffffff",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ConvertToCircle(test.shape)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestGetTypeCircle(t *testing.T) {
	tests := []struct {
		name      string
		shapeType TypeEnum
		want      string
	}{
		{
			name:      "when happy",
			shapeType: CIRCLE,
			want:      string(CIRCLE),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			shape := Circle{
				Type: test.shapeType,
			}
			got := shape.GetType()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestGetAreaCircle(t *testing.T) {
	tests := []struct {
		name   string
		radius float64
		want   float64
	}{
		{
			name:   "when happy",
			radius: 10,
			want:   314.1592653589793,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			shape := Circle{
				Radius: test.radius,
			}
			got := shape.GetArea()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestGetPerimeterCircle(t *testing.T) {
	tests := []struct {
		name   string
		radius float64
		want   float64
	}{
		{
			name:   "when happy",
			radius: 10,
			want:   62.83185307179586,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			shape := Circle{
				Radius: test.radius,
			}
			got := shape.GetPerimeter()
			assert.Equal(t, test.want, got)
		})
	}
}
