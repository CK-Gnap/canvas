package models

import (
	models_interfaces "canvas/models/Interfaces"
	"reflect"
	"testing"
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
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("ConvertToRectangle() = %v, want %v", got, test.want)
			}
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
			if got != test.want {
				t.Errorf("shape.GetTypeRectangle() = %v, want %v", got, test.want)
			}
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
			if got != test.want {
				t.Errorf("shape.GetAreaRectangle() = %v, want %v", got, test.want)
			}
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
			if got != test.want {
				t.Errorf("shape.GetPerimeterRectangle() = %v, want %v", got, test.want)
			}
		})
	}
}
