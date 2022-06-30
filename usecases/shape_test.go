package usecases

import (
	"canvas/constants"
	"canvas/models"
	models_interfaces "canvas/models/Interfaces"
	"canvas/repositories/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateRectangleShapeSuccess(t *testing.T) {

	tests := []struct {
		name     string
		canvasId string
		request  *models.Rectangle
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
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockCanvasRepo := new(mocks.CanvasRepoInterface)
			mockShapeRepo := new(mocks.ShapeRepoInterface)
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
			mockShapeRepo.On("CreateShape", &shape, test.canvasId).Return(nil).Once()

			usecase := NewShapeUsecase(mockCanvasRepo, mockShapeRepo)
			_, err := usecase.CreateRectangleShape(test.request, test.canvasId)

			assert.NoError(t, err)

			mockCanvasRepo.AssertExpectations(t)
			mockShapeRepo.AssertExpectations(t)
		})
	}
}

func TestCreateRectangleShapeError(t *testing.T) {

	tests := []struct {
		name     string
		canvasId string
		request  *models.Rectangle
		errMsg   error
	}{
		{
			name:     "when unhappy",
			canvasId: "1",
			request:  &models.Rectangle{},
			errMsg:   constants.ErrCreateShape,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockCanvasRepo := new(mocks.CanvasRepoInterface)
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			mockCanvasRepo.On("GetCanvas", &models.Canvas{}, test.canvasId).Return(nil).Once()
			mockShapeRepo.On("CreateShape", mock.Anything, test.canvasId).Return(constants.ErrCreateShape).Once()

			usecase := NewShapeUsecase(mockCanvasRepo, mockShapeRepo)
			_, err := usecase.CreateRectangleShape(test.request, test.canvasId)

			assert.Error(t, err)
			assert.Equal(t, test.errMsg, err)

			mockCanvasRepo.AssertExpectations(t)

		})
	}
}

func TestCreateCircleShapeSuccess(t *testing.T) {

	tests := []struct {
		name     string
		canvasId string
		request  *models.Circle
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
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockCanvasRepo := new(mocks.CanvasRepoInterface)
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			mockCanvasRepo.On("GetCanvas", &models.Canvas{}, test.canvasId).Return(nil).Once()
			shape := models.Shape{
				Type:     test.request.Type,
				CanvasId: test.request.CanvasId,
				X:        test.request.X,
				Y:        test.request.Y,
				Radius:   test.request.Radius,
				Color:    test.request.Color,
			}
			mockShapeRepo.On("CreateShape", &shape, test.canvasId).Return(nil).Once()

			usecase := NewShapeUsecase(mockCanvasRepo, mockShapeRepo)
			_, err := usecase.CreateCircleShape(test.request, test.canvasId)

			assert.NoError(t, err)

			mockCanvasRepo.AssertExpectations(t)
			mockShapeRepo.AssertExpectations(t)
		})
	}
}

func TestCreateCircleShapeError(t *testing.T) {

	tests := []struct {
		name     string
		canvasId string
		request  *models.Circle
		errMsg   error
	}{
		{
			name:     "when unhappy",
			canvasId: "1",
			request:  &models.Circle{},
			errMsg:   constants.ErrCreateShape,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockCanvasRepo := new(mocks.CanvasRepoInterface)
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			mockCanvasRepo.On("GetCanvas", &models.Canvas{}, test.canvasId).Return(nil).Once()
			mockShapeRepo.On("CreateShape", mock.Anything, test.canvasId).Return(constants.ErrCreateShape).Once()

			usecase := NewShapeUsecase(mockCanvasRepo, mockShapeRepo)
			_, err := usecase.CreateCircleShape(test.request, test.canvasId)

			assert.Error(t, err)
			assert.Equal(t, test.errMsg, err)

			mockCanvasRepo.AssertExpectations(t)
		})
	}
}

func TestCreateTriangleShapeSuccess(t *testing.T) {

	tests := []struct {
		name     string
		canvasId string
		request  *models.Triangle
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
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockCanvasRepo := new(mocks.CanvasRepoInterface)
			mockShapeRepo := new(mocks.ShapeRepoInterface)
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
			mockShapeRepo.On("CreateShape", &shape, test.canvasId).Return(nil).Once()

			usecase := NewShapeUsecase(mockCanvasRepo, mockShapeRepo)
			_, err := usecase.CreateTriangleShape(test.request, test.canvasId)

			assert.NoError(t, err)

			mockCanvasRepo.AssertExpectations(t)
			mockShapeRepo.AssertExpectations(t)
		})
	}
}

func TestCreateTriangleShapeError(t *testing.T) {

	tests := []struct {
		name     string
		canvasId string
		request  *models.Triangle
		errMsg   error
	}{
		{
			name:     "when unhappy",
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
			errMsg: constants.ErrCreateShape,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockCanvasRepo := new(mocks.CanvasRepoInterface)
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			mockCanvasRepo.On("GetCanvas", &models.Canvas{}, test.canvasId).Return(nil).Once()
			mockShapeRepo.On("CreateShape", mock.Anything, test.canvasId).Return(constants.ErrCreateShape).Once()

			usecase := NewShapeUsecase(mockCanvasRepo, mockShapeRepo)
			_, err := usecase.CreateTriangleShape(test.request, test.canvasId)

			assert.Error(t, err)
			assert.Equal(t, test.errMsg, err)

			mockCanvasRepo.AssertExpectations(t)
		})
	}
}

func TestGetShapesSuccess(t *testing.T) {
	var shapes []models.Shape
	tests := []struct {
		name     string
		canvasId string
		canvas   *models.Canvas
		shapes   *[]models.Shape
	}{
		{
			name:     "when happy",
			canvasId: "1",
			canvas:   &models.Canvas{},
			shapes:   &shapes,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCanvasRepo := new(mocks.CanvasRepoInterface)
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			mockCanvasRepo.On("GetCanvas", test.canvas, test.canvasId).Return(nil).Once()
			mockShapeRepo.On("GetShapes", test.shapes, test.canvasId).Return(test.shapes, nil).Once()

			usecase := NewShapeUsecase(mockCanvasRepo, mockShapeRepo)
			_, err := usecase.GetShapes(test.canvasId)

			assert.NoError(t, err)

			mockCanvasRepo.AssertExpectations(t)
			mockShapeRepo.AssertExpectations(t)
		})
	}
}

func TestGetShapesError(t *testing.T) {
	tests := []struct {
		name     string
		canvasId string
		errMsg   error
	}{
		{
			name:     "when unhappy",
			canvasId: "1",
			errMsg:   constants.ErrGetShapes,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCanvasRepo := new(mocks.CanvasRepoInterface)
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			mockCanvasRepo.On("GetCanvas", &models.Canvas{}, test.canvasId).Return(nil).Once()
			mockShapeRepo.On("GetShapes", mock.Anything, test.canvasId).Return(nil, constants.ErrGetShapes).Once()

			usecase := NewShapeUsecase(mockCanvasRepo, mockShapeRepo)
			_, err := usecase.GetShapes(test.canvasId)

			assert.Error(t, err)
			assert.Equal(t, test.errMsg, err)

			mockCanvasRepo.AssertExpectations(t)
			mockShapeRepo.AssertExpectations(t)
		})
	}
}

func TestGetShapeSuccess(t *testing.T) {

	tests := []struct {
		name  string
		id    string
		shape *models.Shape
	}{
		{
			name:  "when happy",
			id:    "1",
			shape: &models.Shape{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			mockShapeRepo.On("GetShape", test.shape, test.id).Return(nil).Once()
			usecase := NewShapeUsecase(nil, mockShapeRepo)
			_, err := usecase.GetShape(test.shape, test.id)
			assert.NoError(t, err)
			mockShapeRepo.AssertExpectations(t)
		})
	}
}

func TestGetShapeError(t *testing.T) {

	tests := []struct {
		name   string
		id     string
		shape  *models.Shape
		errMsg error
	}{
		{
			name:   "when unhappy",
			id:     "",
			shape:  &models.Shape{},
			errMsg: constants.ErrGetShape,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			mockShapeRepo.On("GetShape", test.shape, test.id).Return(constants.ErrGetShape).Once()

			usecase := NewShapeUsecase(nil, mockShapeRepo)
			_, err := usecase.GetShape(test.shape, test.id)

			assert.Error(t, err)
			assert.Equal(t, test.errMsg, err)

			mockShapeRepo.AssertExpectations(t)
		})
	}
}

func TestUpdateShapeSuccess(t *testing.T) {

	tests := []struct {
		name    string
		id      string
		request *models.Shape
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
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockShapeRepo := new(mocks.ShapeRepoInterface)
			mockShapeRepo.On("GetShape", &models.Shape{}, test.id).Return(nil).Once()
			mockShapeRepo.On("UpdateShape", test.request, test.id).Return(test.request, nil).Once()

			usecase := NewShapeUsecase(nil, mockShapeRepo)
			_, err := usecase.UpdateShape(test.request, test.id)

			assert.NoError(t, err)

			mockShapeRepo.AssertExpectations(t)
		})
	}
}

func TestUpdateShapeError(t *testing.T) {

	tests := []struct {
		name    string
		id      string
		request *models.Shape
		errMsg  error
	}{
		{
			name: "when unhappy",
			id:   "1",
			request: &models.Shape{
				Id:       1,
				Type:     models.CIRCLE,
				CanvasId: 1,
				X:        10,
				Y:        10,
				Radius:   0,
				Color:    "#ffffff",
			},
			errMsg: constants.ErrUpdateShape,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockShapeRepo := new(mocks.ShapeRepoInterface)
			mockShapeRepo.On("GetShape", &models.Shape{}, test.id).Return(nil).Once()
			mockShapeRepo.On("UpdateShape", mock.Anything, test.id).Return(nil, constants.ErrUpdateShape).Once()

			usecase := NewShapeUsecase(nil, mockShapeRepo)
			_, err := usecase.UpdateShape(test.request, test.id)

			assert.Error(t, err)
			assert.Equal(t, test.errMsg, err)

			mockShapeRepo.AssertExpectations(t)
		})
	}
}

func TestDeleteShapeSuccess(t *testing.T) {
	tests := []struct {
		name  string
		id    string
		shape *models.Shape
	}{
		{
			name:  "when happy",
			id:    "1",
			shape: &models.Shape{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			mockShapeRepo.On("GetShape", test.shape, test.id).Return(nil).Once()
			mockShapeRepo.On("DeleteShape", test.shape, test.id).Return(nil).Once()

			usecase := NewShapeUsecase(nil, mockShapeRepo)
			err := usecase.DeleteShape(test.shape, test.id)

			assert.NoError(t, err)

			mockShapeRepo.AssertExpectations(t)
		})
	}
}

func TestDeleteShapeError(t *testing.T) {
	tests := []struct {
		name   string
		id     string
		shape  *models.Shape
		errMsg error
	}{
		{
			name:   "when unhappy",
			id:     "1",
			shape:  &models.Shape{},
			errMsg: constants.ErrDeleteShape,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			mockShapeRepo.On("GetShape", test.shape, test.id).Return(nil).Once()
			mockShapeRepo.On("DeleteShape", mock.Anything, test.id).Return(constants.ErrDeleteShape).Once()

			usecase := NewShapeUsecase(nil, mockShapeRepo)
			err := usecase.DeleteShape(test.shape, test.id)

			assert.Error(t, err)
			assert.Equal(t, test.errMsg, err)

			mockShapeRepo.AssertExpectations(t)
		})
	}
}

func TestConvertTypeOfShape(t *testing.T) {
	tests := []struct {
		name  string
		shape *models.Shape
		want  models_interfaces.ShapeInterface
	}{
		{
			name: "when happy with circle",
			shape: &models.Shape{
				Id:       1,
				Type:     models.CIRCLE,
				CanvasId: 1,
				X:        10,
				Y:        10,
				Radius:   100,
				Color:    "#ffffff",
			},
			want: &models.Circle{
				Id:       1,
				Type:     models.CIRCLE,
				CanvasId: 1,
				X:        10,
				Y:        10,
				Radius:   100,
				Color:    "#ffffff",
			},
		},
		{
			name: "when happy with rectangle",
			shape: &models.Shape{
				Id:       1,
				Type:     models.RECTANGLE,
				CanvasId: 1,
				X:        10,
				Y:        10,
				Width:    100,
				Height:   100,
				Color:    "#ffffff",
			},
			want: &models.Rectangle{
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
			name: "when happy with triangle",
			shape: &models.Shape{
				Id:       1,
				Type:     models.TRIANGLE,
				CanvasId: 1,
				X:        10,
				Y:        10,
				Width:    100,
				Height:   100,
				Color:    "#ffffff",
			},
			want: &models.Triangle{
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
			name: "when unhappy with unknown shape",
			shape: &models.Shape{
				Id:       1,
				Type:     "unknown",
				CanvasId: 1,
				X:        10,
				Y:        10,
				Width:    100,
				Height:   100,
				Color:    "#ffffff",
			},
			want: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ConvertTypeOfShape(test.shape)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestConvertTypeOfShapes(t *testing.T) {
	tests := []struct {
		name   string
		shapes *[]models.Shape
		want   []models_interfaces.ShapeInterface
	}{
		{
			name: "when happy",
			shapes: &[]models.Shape{
				{
					Id:       1,
					Type:     models.CIRCLE,
					CanvasId: 1,
					X:        10,
					Y:        10,
					Radius:   100,
					Color:    "#ffffff",
				},
				{
					Id:       2,
					Type:     models.RECTANGLE,
					CanvasId: 1,
					X:        10,
					Y:        10,
					Width:    100,
					Height:   100,
					Color:    "#ffffff",
				},
				{
					Id:       3,
					Type:     models.TRIANGLE,
					CanvasId: 1,
					X:        10,
					Y:        10,
					Width:    100,
					Height:   100,
					Color:    "#ffffff",
				},
			},
			want: []models_interfaces.ShapeInterface{
				&models.Circle{
					Id:       1,
					Type:     models.CIRCLE,
					CanvasId: 1,
					X:        10,
					Y:        10,
					Radius:   100,
					Color:    "#ffffff",
				},
				&models.Rectangle{
					Id:       2,
					Type:     models.RECTANGLE,
					CanvasId: 1,
					X:        10,
					Y:        10,
					Width:    100,
					Height:   100,
					Color:    "#ffffff",
				},
				&models.Triangle{
					Id:       3,
					Type:     models.TRIANGLE,
					CanvasId: 1,
					X:        10,
					Y:        10,
					Width:    100,
					Height:   100,
					Color:    "#ffffff",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ConvertTypeOfShapes(test.shapes)
			assert.Equal(t, test.want, got)
		})
	}
}
