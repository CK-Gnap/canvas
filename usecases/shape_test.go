package usecases

import (
	"canvas/models"
	"canvas/repositories/Interfaces/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRectangleShape(t *testing.T) {

	tests := []struct {
		name     string
		canvasId string
		request  *models.Rectangle
		isErr    bool
	}{
		{
			name:     "when happy",
			canvasId: "1",
			request: &models.Rectangle{
				CanvasId: 1,
				Type:     models.RECTANGLE,
				X:        10,
				Y:        10,
				Width:    100,
				Height:   100,
				Color:    "#ffffff",
			},
			isErr: false,
		},
		{
			name:     "when unhappy",
			canvasId: "",
			request: &models.Rectangle{
				CanvasId: 1,
				Type:     models.RECTANGLE,
				X:        10,
				Y:        10,
				Width:    0,
				Height:   100,
				Color:    "#ffffff",
			},
			isErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockCanvasRepo := new(mocks.CanvasRepoInterface)
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			if test.canvasId == "" {
				mockCanvasRepo.On("GetCanvas", &models.Canvas{}, test.canvasId).Return(errors.New("Error getting canvas")).Once()
				usecase := NewCanvasUsecase(mockCanvasRepo, mockShapeRepo)
				_, err := usecase.GetCanvas(&models.Canvas{}, test.canvasId)
				assert.Error(t, err)
				mockCanvasRepo.AssertExpectations(t)
			} else {
				mockCanvasRepo.On("GetCanvas", &models.Canvas{}, test.canvasId).Return(nil).Once()
				shape := models.Shape{
					Type:     test.request.Type,
					CanvasId: test.request.CanvasId,
					X:        test.request.X,
					Y:        test.request.Y,
					Width:    test.request.Width,
					Height:   test.request.Height,
					Color:    test.request.Color,
				}
				if test.isErr {
					mockShapeRepo.On("CreateShape", &shape, test.canvasId).Return(errors.New("Error creating shape")).Once()
				} else {
					mockShapeRepo.On("CreateShape", &shape, test.canvasId).Return(nil).Once()
				}

				usecase := NewShapeUsecase(mockCanvasRepo, mockShapeRepo)
				_, err := usecase.CreateRectangleShape(test.request, test.canvasId)
				if test.isErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}
				mockCanvasRepo.AssertExpectations(t)
				mockShapeRepo.AssertExpectations(t)
			}
		})
	}
}

func TestCreateCircleShape(t *testing.T) {

	tests := []struct {
		name     string
		canvasId string
		request  *models.Circle
		isErr    bool
	}{
		{
			name:     "when happy",
			canvasId: "1",
			request: &models.Circle{
				CanvasId: 1,
				Type:     models.CIRCLE,
				X:        10,
				Y:        10,
				Radius:   100,
				Color:    "#ffffff",
			},
			isErr: false,
		},
		{
			name:     "when unhappy",
			canvasId: "",
			request: &models.Circle{
				CanvasId: 1,
				Type:     models.CIRCLE,
				X:        10,
				Y:        10,
				Radius:   0,
				Color:    "#ffffff",
			},
			isErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockCanvasRepo := new(mocks.CanvasRepoInterface)
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			if test.canvasId == "" {
				mockCanvasRepo.On("GetCanvas", &models.Canvas{}, test.canvasId).Return(errors.New("Error getting canvas")).Once()
				usecase := NewCanvasUsecase(mockCanvasRepo, mockShapeRepo)
				_, err := usecase.GetCanvas(&models.Canvas{}, test.canvasId)
				assert.Error(t, err)
				mockCanvasRepo.AssertExpectations(t)
			} else {
				mockCanvasRepo.On("GetCanvas", &models.Canvas{}, test.canvasId).Return(nil).Once()
				shape := models.Shape{
					Type:     test.request.Type,
					CanvasId: test.request.CanvasId,
					X:        test.request.X,
					Y:        test.request.Y,
					Radius:   test.request.Radius,
					Color:    test.request.Color,
				}
				if test.isErr {
					mockShapeRepo.On("CreateShape", &shape, test.canvasId).Return(errors.New("Error creating shape")).Once()
				} else {
					mockShapeRepo.On("CreateShape", &shape, test.canvasId).Return(nil).Once()
				}

				usecase := NewShapeUsecase(mockCanvasRepo, mockShapeRepo)
				_, err := usecase.CreateCircleShape(test.request, test.canvasId)
				if test.isErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}
				mockCanvasRepo.AssertExpectations(t)
				mockShapeRepo.AssertExpectations(t)
			}
		})
	}
}

func TestCreateTriangleShape(t *testing.T) {

	tests := []struct {
		name     string
		canvasId string
		request  *models.Triangle
		isErr    bool
	}{
		{
			name:     "when happy",
			canvasId: "1",
			request: &models.Triangle{
				CanvasId: 1,
				Type:     models.TRIANGLE,
				X:        10,
				Y:        10,
				Width:    100,
				Height:   100,
				Color:    "#ffffff",
			},
			isErr: false,
		},
		{
			name:     "when unhappy",
			canvasId: "",
			request: &models.Triangle{
				CanvasId: 1,
				Type:     models.TRIANGLE,
				X:        10,
				Y:        10,
				Width:    0,
				Height:   0,
				Color:    "#ffffff",
			},
			isErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockCanvasRepo := new(mocks.CanvasRepoInterface)
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			if test.canvasId == "" {
				mockCanvasRepo.On("GetCanvas", &models.Canvas{}, test.canvasId).Return(errors.New("Error getting canvas")).Once()
				usecase := NewCanvasUsecase(mockCanvasRepo, mockShapeRepo)
				_, err := usecase.GetCanvas(&models.Canvas{}, test.canvasId)
				assert.Error(t, err)
				mockCanvasRepo.AssertExpectations(t)
			} else {
				mockCanvasRepo.On("GetCanvas", &models.Canvas{}, test.canvasId).Return(nil).Once()
				shape := models.Shape{
					Type:     test.request.Type,
					CanvasId: test.request.CanvasId,
					X:        test.request.X,
					Y:        test.request.Y,
					Width:    test.request.Width,
					Height:   test.request.Height,
					Color:    test.request.Color,
				}
				if test.isErr {
					mockShapeRepo.On("CreateShape", &shape, test.canvasId).Return(errors.New("Error creating shape")).Once()
				} else {
					mockShapeRepo.On("CreateShape", &shape, test.canvasId).Return(nil).Once()
				}

				usecase := NewShapeUsecase(mockCanvasRepo, mockShapeRepo)
				_, err := usecase.CreateTriangleShape(test.request, test.canvasId)
				if test.isErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}
				mockCanvasRepo.AssertExpectations(t)
				mockShapeRepo.AssertExpectations(t)
			}
		})
	}
}

// test get shapes usecase
// func TestGetShapes(t *testing.T) {
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
// 		name     string
// 		canvasId string
// 		isErr    bool
// 	}{
// 		{
// 			name:     "when happy",
// 			canvasId: "1",
// 			isErr:    false,
// 		},
// 	}
// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			mockShapeRepo := new(mocks.ShapeRepoInterface)
// 			if test.isErr {
// 				mockShapeRepo.On("GetShapes", &shapes, test.canvasId).Return(&shapes, errors.New("Error getting shapes")).Once()
// 			} else {
// 				mockShapeRepo.On("GetShapes", &shapes, test.canvasId).Return(&shapes, nil).Once()
// 			}
// 			usecase := NewShapeUsecase(nil, mockShapeRepo)
// 			_, err := usecase.GetShapes(test.canvasId)
// 			if test.isErr {
// 				assert.Error(t, err)
// 			} else {
// 				assert.NoError(t, err)
// 			}
// 			mockShapeRepo.AssertExpectations(t)
// 		})
// 	}
// }

func TestGetShape(t *testing.T) {

	tests := []struct {
		name     string
		id       string
		idString string
		isErr    bool
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
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			if test.id == "" {
				mockShapeRepo.On("GetShape", &models.Shape{}, test.id).Return(errors.New("Error getting shape")).Once()
				usecase := NewShapeUsecase(nil, mockShapeRepo)
				_, err := usecase.GetShape(&models.Shape{}, test.id)
				assert.Error(t, err)
				mockShapeRepo.AssertExpectations(t)
			} else {
				mockShapeRepo.On("GetShape", &models.Shape{}, test.id).Return(nil).Once()
				usecase := NewShapeUsecase(nil, mockShapeRepo)
				_, err := usecase.GetShape(&models.Shape{}, test.id)
				if test.isErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}
				mockShapeRepo.AssertExpectations(t)
			}
		})
	}
}

func TestUpdateShape(t *testing.T) {

	tests := []struct {
		name    string
		id      string
		request *models.Shape
		isErr   bool
	}{
		{
			name: "when happy with rectangle",
			id:   "1",
			request: &models.Shape{
				Id:       1,
				Type:     models.RECTANGLE,
				CanvasId: 1,
				X:        10,
				Y:        10,
				Width:    100,
				Height:   100,
				Color:    "#ffffff",
			},
		},
		{
			name: "when happy with circle",
			id:   "1",
			request: &models.Shape{
				Id:       1,
				Type:     models.CIRCLE,
				CanvasId: 1,
				X:        10,
				Y:        10,
				Radius:   100,
				Color:    "#ffffff",
			},
			isErr: false,
		},
		{
			name: "when happy with triangle",
			id:   "1",
			request: &models.Shape{
				Id:       1,
				Type:     models.TRIANGLE,
				CanvasId: 1,
				X:        10,
				Y:        10,
				Width:    100,
				Height:   100,
				Color:    "#ffffff",
			},
		},
		{
			name: "when unhappy",
			id:   "",
			request: &models.Shape{
				Id:       1,
				Type:     models.CIRCLE,
				CanvasId: 1,
				X:        10,
				Y:        10,
				Radius:   0,
				Color:    "#ffffff",
			},
			isErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockShapeRepo := new(mocks.ShapeRepoInterface)
			if test.id == "" {
				mockShapeRepo.On("GetShape", &models.Shape{}, test.id).Return(errors.New("Error getting shape")).Once()
				usecase := NewShapeUsecase(nil, mockShapeRepo)
				_, err := usecase.GetShape(&models.Shape{}, test.id)
				assert.Error(t, err)
				mockShapeRepo.AssertExpectations(t)
			} else {
				mockShapeRepo.On("GetShape", &models.Shape{}, test.id).Return(nil).Once()
				if test.isErr {
					mockShapeRepo.On("UpdateShape", test.request, test.id).Return(nil, errors.New("Error updating shape")).Once()
				} else {
					mockShapeRepo.On("UpdateShape", test.request, test.id).Return(test.request, nil).Once()
				}
				usecase := NewShapeUsecase(nil, mockShapeRepo)
				_, err := usecase.UpdateShape(test.request, test.id)
				if test.isErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}
				mockShapeRepo.AssertExpectations(t)
			}
		})
	}
}

func TestDeleteShape(t *testing.T) {
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
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			if test.id == "" {
				mockShapeRepo.On("GetShape", &models.Shape{}, test.id).Return(errors.New("Error getting shape")).Once()
				usecase := NewShapeUsecase(nil, mockShapeRepo)
				_, err := usecase.GetShape(&models.Shape{}, test.id)
				assert.Error(t, err)
				mockShapeRepo.AssertExpectations(t)
			} else {
				mockShapeRepo.On("GetShape", &models.Shape{}, test.id).Return(nil).Once()
				mockShapeRepo.On("DeleteShape", &models.Shape{}, test.id).Return(nil).Once()
				usecase := NewShapeUsecase(nil, mockShapeRepo)
				err := usecase.DeleteShape(&models.Shape{}, test.id)
				if test.isErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}
				mockShapeRepo.AssertExpectations(t)
			}
		})
	}
}
