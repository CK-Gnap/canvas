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
			if test.request.Name == "" {
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
			if test.request.Name == "" {
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
			if test.id == "" {
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

// func TestGetCanvas(t *testing.T) {
// 	shapes := []models.Shape{
// 		{
// 			Id:       1,
// 			Type:     models.CIRCLE,
// 			CanvasId: 1,
// 			X:        10,
// 			Y:        10,
// 			Radius:   100,
// 			Color:    "#ffffff",
// 		},
// 		{
// 			Id:       2,
// 			Type:     models.TRIANGLE,
// 			CanvasId: 1,
// 			X:        10,
// 			Y:        10,
// 			Width:    100,
// 			Height:   100,
// 			Color:    "#ffffff",
// 		},
// 		{
// 			Id:       3,
// 			Type:     models.RECTANGLE,
// 			CanvasId: 1,
// 			X:        10,
// 			Y:        10,
// 			Width:    100,
// 			Height:   100,
// 			Color:    "#ffffff",
// 		},
// 	}
// 	tests := []struct {
// 		name  string
// 		id    string
// 		isErr bool
// 	}{
// 		{
// 			name:  "when happy",
// 			id:    "1",
// 			isErr: false,
// 		},
// 		{
// 			name:  "when unhappy",
// 			id:    "",
// 			isErr: true,
// 		},
// 	}
// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {

// 			mockCanvasRepo := new(mocks.CanvasRepoInterface)
// 			mockShapeRepo := new(mocks.ShapeRepoInterface)
// 			if test.id == "" {
// 				mockCanvasRepo.On("GetCanvas", &models.Canvas{}, test.id).Return(errors.New("Error getting canvas")).Once()
// 			} else {
// 				mockCanvasRepo.On("GetCanvas", &models.Canvas{}, test.id).Return(nil).Once()
// 			}

// 			mockShapeRepo.On("GetShapes", shapes, test.id).Return(shapes, nil).Once()

// 			usecase := NewCanvasUsecase(mockCanvasRepo, mockShapeRepo)
// 			canvas, err := usecase.GetCanvas(&models.Canvas{}, test.id)
// 			if test.isErr {
// 				assert.Error(t, err)
// 			} else {
// 				assert.NoError(t, err)
// 				assert.NotNil(t, canvas)
// 			}
// 			mockCanvasRepo.AssertExpectations(t)
// 		})
// 	}
// }
