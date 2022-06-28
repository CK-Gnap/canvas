package models

import (
	models_interfaces "canvas/models/Interfaces"
	"reflect"
	"testing"
)

func TestConvertToTriangle(t *testing.T) {
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
				Type:     TRIANGLE,
				X:        1,
				Y:        1,
				Width:    1,
				Height:   1,
				Color:    "#000000",
			},
			want: &Triangle{
				Id:       1,
				CanvasId: 1,
				Type:     TRIANGLE,
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
			got := ConvertToTriangle(test.shape)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("ConvertToTriangle() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestGetTypeTriangle(t *testing.T) {
	tests := []struct {
		name      string
		shapeType TypeEnum
		want      string
	}{
		{
			name:      "when happy",
			shapeType: TRIANGLE,
			want:      string(TRIANGLE),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			shape := Triangle{
				Type: test.shapeType,
			}
			got := shape.GetType()
			if got != test.want {
				t.Errorf("shape.GetTypeTriangle() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestGetAreaTriangle(t *testing.T) {
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
			want:   50,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			shape := Triangle{
				Width:  test.Width,
				Height: test.Height,
			}
			got := shape.GetArea()
			if got != test.want {
				t.Errorf("shape.GetAreaTriangle() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestGetPerimeterTriangle(t *testing.T) {
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
			want:   30,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			shape := Triangle{
				Width:  test.Width,
				Height: test.Height,
			}
			got := shape.GetPerimeter()
			if got != test.want {
				t.Errorf("shape.GetPerimeterTriangle() = %v, want %v", got, test.want)
			}
		})
	}
}
