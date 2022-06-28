package usecases

import (
	models "canvas/models"

	"canvas/repositories/Interfaces/mocks"
	"errors"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCanvas(t *testing.T) {

	tests := []struct {
		name    string
		request *models.Canvas
		isErr   bool
	}{
		{
			name: "when happy",
			request: &models.Canvas{
				Name:   "test",
				Width:  100,
				Height: 100,
				Color:  "#ffffff",
			},
			isErr: false,
		},
		{
			name: "when unhappy",
			request: &models.Canvas{
				Name:   "",
				Width:  100,
				Height: 100,
				Color:  "#ffffff",
			},
			isErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockCanvasRepo := new(mocks.CanvasRepoInterface)
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			if test.isErr {
				mockCanvasRepo.On("CreateCanvas", test.request).Return(errors.New("Error creating canvas")).Once()
			} else {
				mockCanvasRepo.On("CreateCanvas", test.request).Return(nil).Once()
			}
			usecase := NewCanvasUsecase(mockCanvasRepo, mockShapeRepo)
			canvas, err := usecase.CreateCanvas(test.request)
			if test.isErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, canvas)
			}

			mockCanvasRepo.AssertExpectations(t)
		})
	}
}

func TestGetCanvases(t *testing.T) {
	var canvases []models.Canvas

	tests := []struct {
		name     string
		canvases *[]models.Canvas
		isErr    bool
	}{
		{
			name: "when happy",

			isErr: false,
		},
		{
			name:  "when unhappy",
			isErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockCanvasRepo := new(mocks.CanvasRepoInterface)
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			if test.isErr {
				mockCanvasRepo.On("GetCanvases", &canvases).Return(errors.New("Error getting canvases")).Once()
			} else {
				mockCanvasRepo.On("GetCanvases", &canvases).Return(nil).Once()
			}

			usecaseCanvas := NewCanvasUsecase(mockCanvasRepo, mockShapeRepo)
			canvas, err := usecaseCanvas.GetCanvases()
			if test.isErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, &canvases, &canvas)
			}
			mockCanvasRepo.AssertExpectations(t)
		})
	}
}

func TestGetCanvas(t *testing.T) {
	var shapes []models.Shape
	tests := []struct {
		name  string
		id    string
		isErr bool
	}{
		{
			name:  "when happy",
			id:    "1",
			isErr: false,
		},
		{
			name:  "when unhappy",
			id:    "",
			isErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockCanvasRepo := new(mocks.CanvasRepoInterface)
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			if test.isErr {
				mockCanvasRepo.On("GetCanvas", &models.Canvas{}, test.id).Return(errors.New("Error getting canvas")).Once()
			} else {
				mockCanvasRepo.On("GetCanvas", &models.Canvas{}, test.id).Return(nil).Once()
			}

			mockShapeRepo.On("GetShapes", &shapes, "0").Return(&shapes, nil).Once()

			usecaseCanvas := NewCanvasUsecase(mockCanvasRepo, mockShapeRepo)
			canvas, errCanvas := usecaseCanvas.GetCanvas(&models.Canvas{}, test.id)
			if test.isErr {
				assert.Error(t, errCanvas)
			} else {
				assert.NoError(t, errCanvas)
				assert.NotNil(t, canvas)
			}
			mockCanvasRepo.AssertExpectations(t)
			// mockShapeRepo.AssertExpectations(t)
		})
	}
}

func TestUpdateCanvas(t *testing.T) {

	tests := []struct {
		name    string
		id      string
		request *models.Canvas
		isErr   bool
	}{
		{
			name: "when happy",
			id:   "1",
			request: &models.Canvas{
				Name:   "test",
				Width:  100,
				Height: 100,
				Color:  "#ffffff",
			},
			isErr: false,
		},
		{
			name: "when unhappy",
			id:   "1",
			request: &models.Canvas{
				Name:   "",
				Width:  100,
				Height: 100,
				Color:  "#ffffff",
			},
			isErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockCanvasRepo := new(mocks.CanvasRepoInterface)
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			mockCanvasRepo.On("GetCanvas", &models.Canvas{}, test.id).Return(nil).Once()
			if test.isErr {
				mockCanvasRepo.On("UpdateCanvas", test.request, test.id).Return(errors.New("Error updating canvas")).Once()
			} else {
				mockCanvasRepo.On("UpdateCanvas", test.request, test.id).Return(nil).Once()
			}
			usecase := NewCanvasUsecase(mockCanvasRepo, mockShapeRepo)
			canvas, err := usecase.UpdateCanvas(test.request, test.id)
			if test.isErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, canvas)
			}
			mockCanvasRepo.AssertExpectations(t)
		})
	}
}

func TestDeleteCanvas(t *testing.T) {
	canvas := &models.Canvas{
		Id:     1,
		Name:   "test",
		Width:  100,
		Height: 100,
		Color:  "#ffffff",
	}
	tests := []struct {
		name  string
		id    string
		isErr bool
	}{
		{
			name:  "when happy",
			id:    "1",
			isErr: false,
		},
		{
			name:  "when unhappy",
			id:    "",
			isErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockCanvasRepo := new(mocks.CanvasRepoInterface)
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			mockCanvasRepo.On("GetCanvas", &models.Canvas{}, test.id).Return(nil).Once()
			if test.isErr {
				mockCanvasRepo.On("DeleteCanvas", canvas, test.id).Return(errors.New("Error deleting canvas")).Once()
			} else {
				mockCanvasRepo.On("DeleteCanvas", canvas, test.id).Return(nil).Once()
			}
			usecase := NewCanvasUsecase(mockCanvasRepo, mockShapeRepo)
			err := usecase.DeleteCanvas(canvas, test.id)
			if test.isErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			mockCanvasRepo.AssertExpectations(t)
		})
	}
}

func TestGetTotalArea(t *testing.T) {
	var shapes []models.Shape
	tests := []struct {
		name   string
		id     string
		canvas models.Canvas
		want   float64
	}{
		{
			name: "when happy",
			id:   "1",
			canvas: models.Canvas{
				Id:     1,
				Name:   "test",
				Width:  100,
				Height: 100,
				Color:  "#ffffff",
			},
			want: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCanvasRepo := new(mocks.CanvasRepoInterface)
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			mockCanvasRepo.On("GetCanvas", &test.canvas, test.id).Return(nil).Once()
			mockShapeRepo.On("GetShapes", &shapes, test.id).Return(&shapes, nil).Once()
			usecase := NewCanvasUsecase(mockCanvasRepo, mockShapeRepo)
			got, _ := usecase.GetTotalArea(&test.canvas, test.id)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
			mockCanvasRepo.AssertExpectations(t)
			mockShapeRepo.AssertExpectations(t)
		})
	}
}

func TestGetTotalPerimeter(t *testing.T) {
	var shapes []models.Shape
	tests := []struct {
		name   string
		id     string
		canvas models.Canvas
		want   float64
	}{
		{
			name: "when happy",
			id:   "1",
			canvas: models.Canvas{
				Id:     1,
				Name:   "test",
				Width:  100,
				Height: 100,
				Color:  "#ffffff",
			},
			want: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCanvasRepo := new(mocks.CanvasRepoInterface)
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			mockCanvasRepo.On("GetCanvas", &test.canvas, test.id).Return(nil).Once()
			mockShapeRepo.On("GetShapes", &shapes, test.id).Return(&shapes, nil).Once()
			usecase := NewCanvasUsecase(mockCanvasRepo, mockShapeRepo)
			got, _ := usecase.GetTotalPerimeter(&test.canvas, test.id)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
			mockCanvasRepo.AssertExpectations(t)
			mockShapeRepo.AssertExpectations(t)
		})
	}
}

func TestDrawCanvas(t *testing.T) {
	tests := []struct {
		name   string
		id     string
		canvas models.Canvas
		want   string
	}{
		{
			name: "when happy",
			id:   "1",
			canvas: models.Canvas{
				Id:     1,
				Name:   "test",
				Width:  100,
				Height: 100,
				Color:  "#ffffff",
			},
			want: "test.jpg",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			usecase := CanvasUsecase{}
			got, _ := usecase.DrawCanvas(&test.canvas, test.id)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
