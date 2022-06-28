package models

import (
	models_interfaces "canvas/models/Interfaces"
	"reflect"
	"testing"
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
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("ConvertToCircle() = %v, want %v", got, test.want)
			}
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
			if got != test.want {
				t.Errorf("shape.GetTypeCircle() = %v, want %v", got, test.want)
			}
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
			if got != test.want {
				t.Errorf("shape.GetAreaCircle() = %v, want %v", got, test.want)
			}
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
			if got != test.want {
				t.Errorf("shape.GetPerimeterCircle() = %v, want %v", got, test.want)
			}
		})
	}
}
