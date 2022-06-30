package usecases

import (
	"canvas/constants"
	models "canvas/models"

	"canvas/repositories/mocks"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCanvasSuccess(t *testing.T) {

	tests := []struct {
		name    string
		request *models.Canvas
	}{
		{
			name: "when happy",
			request: &models.Canvas{
				Name:   "test",
				Width:  100,
				Height: 100,
				Color:  "#ffffff",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockCanvasRepo := new(mocks.CanvasRepoInterface)
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			mockCanvasRepo.On("CreateCanvas", test.request).Return(nil).Once()

			usecase := NewCanvasUsecase(mockCanvasRepo, mockShapeRepo)
			canvas, err := usecase.CreateCanvas(test.request)

			assert.NoError(t, err)
			assert.NotNil(t, canvas)

			mockCanvasRepo.AssertExpectations(t)
		})
	}
}

func TestCreateCanvasError(t *testing.T) {
	tests := []struct {
		name    string
		request *models.Canvas
		errMsg  error
	}{
		{
			name: "when unhappy",
			request: &models.Canvas{
				Name:   "",
				Width:  100,
				Height: 100,
				Color:  "#ffffff",
			},
			errMsg: constants.ErrCreateCanvas,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockCanvasRepo := new(mocks.CanvasRepoInterface)
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			mockCanvasRepo.On("CreateCanvas", test.request).Return(constants.ErrCreateCanvas).Once()

			usecase := NewCanvasUsecase(mockCanvasRepo, mockShapeRepo)
			_, err := usecase.CreateCanvas(test.request)

			assert.Error(t, err)
			assert.Equal(t, test.errMsg, err)

			mockCanvasRepo.AssertExpectations(t)
		})
	}
}

func TestGetCanvasesSuccess(t *testing.T) {
	var canvases []models.Canvas

	tests := []struct {
		name     string
		canvases *[]models.Canvas
	}{
		{
			name:     "when happy",
			canvases: &canvases,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockCanvasRepo := new(mocks.CanvasRepoInterface)
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			mockCanvasRepo.On("GetCanvases", test.canvases).Return(nil).Once()

			usecaseCanvas := NewCanvasUsecase(mockCanvasRepo, mockShapeRepo)
			canvas, err := usecaseCanvas.GetCanvases()

			assert.NoError(t, err)
			assert.Equal(t, test.canvases, &canvas)

			mockCanvasRepo.AssertExpectations(t)
		})
	}
}

func TestGetCanvasesError(t *testing.T) {
	tests := []struct {
		name   string
		errMsg error
	}{
		{
			name:   "when unhappy",
			errMsg: constants.ErrGetCanvases,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockCanvasRepo := new(mocks.CanvasRepoInterface)
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			mockCanvasRepo.On("GetCanvases", mock.Anything).Return(constants.ErrGetCanvases).Once()

			usecaseCanvas := NewCanvasUsecase(mockCanvasRepo, mockShapeRepo)
			_, err := usecaseCanvas.GetCanvases()

			assert.Error(t, err)
			assert.Equal(t, test.errMsg, err)

			mockCanvasRepo.AssertExpectations(t)
		})
	}
}

func TestGetCanvasSuccess(t *testing.T) {
	var shapes []models.Shape
	tests := []struct {
		name   string
		id     string
		canvas *models.Canvas
	}{
		{
			name:   "when happy",
			id:     "1",
			canvas: &models.Canvas{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockCanvasRepo := new(mocks.CanvasRepoInterface)
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			mockCanvasRepo.On("GetCanvas", test.canvas, test.id).Return(nil).Once()
			mockShapeRepo.On("GetShapes", &shapes, mock.AnythingOfType("string")).Return(&shapes, nil).Once()

			usecaseCanvas := NewCanvasUsecase(mockCanvasRepo, mockShapeRepo)
			canvas, errCanvas := usecaseCanvas.GetCanvas(test.canvas, test.id)
			assert.NoError(t, errCanvas)
			assert.NotNil(t, canvas)

			mockCanvasRepo.AssertExpectations(t)
			mockShapeRepo.AssertExpectations(t)
		})
	}
}

func TestGetCanvasError(t *testing.T) {
	tests := []struct {
		name   string
		id     string
		canvas *models.Canvas
		errMsg error
	}{
		{
			name:   "when unhappy",
			id:     "",
			canvas: &models.Canvas{},
			errMsg: constants.ErrGetCanvas,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockCanvasRepo := new(mocks.CanvasRepoInterface)
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			mockCanvasRepo.On("GetCanvas", test.canvas, test.id).Return(constants.ErrGetCanvas).Once()

			usecaseCanvas := NewCanvasUsecase(mockCanvasRepo, mockShapeRepo)
			_, errCanvas := usecaseCanvas.GetCanvas(test.canvas, test.id)

			assert.Error(t, errCanvas)
			assert.Equal(t, test.errMsg, errCanvas)

			mockCanvasRepo.AssertExpectations(t)
		})
	}
}

func TestUpdateCanvasSuccess(t *testing.T) {

	tests := []struct {
		name    string
		id      string
		canvas  *models.Canvas
		request *models.Canvas
	}{
		{
			name:   "when happy",
			id:     "1",
			canvas: &models.Canvas{},
			request: &models.Canvas{
				Name:   "test",
				Width:  100,
				Height: 100,
				Color:  "#ffffff",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockCanvasRepo := new(mocks.CanvasRepoInterface)
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			mockCanvasRepo.On("GetCanvas", test.canvas, test.id).Return(nil).Once()
			mockCanvasRepo.On("UpdateCanvas", test.request, test.id).Return(nil).Once()

			usecase := NewCanvasUsecase(mockCanvasRepo, mockShapeRepo)
			canvas, err := usecase.UpdateCanvas(test.request, test.id)

			assert.NoError(t, err)
			assert.NotNil(t, canvas)

			mockCanvasRepo.AssertExpectations(t)
		})
	}
}

func TestUpdateCanvasError(t *testing.T) {

	tests := []struct {
		name    string
		id      string
		canvas  *models.Canvas
		request *models.Canvas
		errMsg  error
	}{
		{
			name:   "when unhappy",
			id:     "1",
			canvas: &models.Canvas{},
			request: &models.Canvas{
				Name:   "",
				Width:  100,
				Height: 100,
				Color:  "#ffffff",
			},
			errMsg: constants.ErrUpdateCanvas,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockCanvasRepo := new(mocks.CanvasRepoInterface)
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			mockCanvasRepo.On("GetCanvas", test.canvas, test.id).Return(nil).Once()
			mockCanvasRepo.On("UpdateCanvas", test.request, test.id).Return(constants.ErrUpdateCanvas).Once()

			usecase := NewCanvasUsecase(mockCanvasRepo, mockShapeRepo)
			_, err := usecase.UpdateCanvas(test.request, test.id)

			assert.Error(t, err)
			assert.Equal(t, test.errMsg, err)

			mockCanvasRepo.AssertExpectations(t)
		})
	}
}

func TestDeleteCanvasSuccess(t *testing.T) {
	tests := []struct {
		name   string
		id     string
		canvas *models.Canvas
	}{
		{
			name:   "when happy",
			id:     "1",
			canvas: &models.Canvas{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockCanvasRepo := new(mocks.CanvasRepoInterface)
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			mockCanvasRepo.On("GetCanvas", test.canvas, test.id).Return(nil).Once()
			mockCanvasRepo.On("DeleteCanvas", test.canvas, test.id).Return(nil).Once()

			usecase := NewCanvasUsecase(mockCanvasRepo, mockShapeRepo)
			err := usecase.DeleteCanvas(test.canvas, test.id)

			assert.NoError(t, err)

			mockCanvasRepo.AssertExpectations(t)
		})
	}
}

func TestDeleteCanvasError(t *testing.T) {
	tests := []struct {
		name   string
		id     string
		canvas *models.Canvas
		errMsg error
	}{
		{
			name:   "when unhappy",
			id:     "",
			canvas: &models.Canvas{},
			errMsg: constants.ErrDeleteCanvas,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockCanvasRepo := new(mocks.CanvasRepoInterface)
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			mockCanvasRepo.On("GetCanvas", test.canvas, test.id).Return(nil).Once()
			mockCanvasRepo.On("DeleteCanvas", test.canvas, test.id).Return(constants.ErrDeleteCanvas).Once()

			usecase := NewCanvasUsecase(mockCanvasRepo, mockShapeRepo)
			err := usecase.DeleteCanvas(test.canvas, test.id)

			assert.Error(t, err)
			assert.Equal(t, test.errMsg, err)

			mockCanvasRepo.AssertExpectations(t)
		})
	}
}

func TestGetTotalAreaSuccess(t *testing.T) {
	var shapes []models.Shape
	tests := []struct {
		name   string
		id     string
		canvas *models.Canvas
	}{
		{
			name: "when happy",
			id:   "1",
			canvas: &models.Canvas{
				Id:     1,
				Name:   "test",
				Width:  100,
				Height: 100,
				Color:  "#ffffff",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCanvasRepo := new(mocks.CanvasRepoInterface)
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			mockCanvasRepo.On("GetCanvas", test.canvas, test.id).Return(nil).Once()
			mockShapeRepo.On("GetShapes", &shapes, test.id).Return(&shapes, nil).Once()

			usecase := NewCanvasUsecase(mockCanvasRepo, mockShapeRepo)
			_, err := usecase.GetTotalArea(test.canvas, test.id)

			assert.NoError(t, err)

			mockCanvasRepo.AssertExpectations(t)
			mockShapeRepo.AssertExpectations(t)
		})
	}
}

func TestGetTotalAreaError(t *testing.T) {
	var shapes []models.Shape
	tests := []struct {
		name   string
		id     string
		canvas *models.Canvas
		want   int
	}{
		{
			name: "when unhappy",
			id:   "1",
			canvas: &models.Canvas{
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
			mockCanvasRepo.On("GetCanvas", test.canvas, test.id).Return(nil).Once()
			mockShapeRepo.On("GetShapes", &shapes, test.id).Return(&shapes, nil).Once()

			usecase := NewCanvasUsecase(mockCanvasRepo, mockShapeRepo)
			total, _ := usecase.GetTotalArea(test.canvas, test.id)

			assert.NotEqual(t, test.want, total)

			mockCanvasRepo.AssertExpectations(t)
			mockShapeRepo.AssertExpectations(t)
		})
	}
}

func TestGetTotalPerimeterSuccess(t *testing.T) {
	var shapes []models.Shape
	tests := []struct {
		name   string
		id     string
		canvas models.Canvas
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
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCanvasRepo := new(mocks.CanvasRepoInterface)
			mockShapeRepo := new(mocks.ShapeRepoInterface)
			mockCanvasRepo.On("GetCanvas", &test.canvas, test.id).Return(nil).Once()
			mockShapeRepo.On("GetShapes", &shapes, test.id).Return(&shapes, nil).Once()

			usecase := NewCanvasUsecase(mockCanvasRepo, mockShapeRepo)
			_, err := usecase.GetTotalPerimeter(&test.canvas, test.id)

			assert.NoError(t, err)

			mockCanvasRepo.AssertExpectations(t)
			mockShapeRepo.AssertExpectations(t)
		})
	}
}

func TestGetTotalPerimeterError(t *testing.T) {
	var shapes []models.Shape
	tests := []struct {
		name   string
		id     string
		canvas models.Canvas
		want   int
	}{
		{
			name: "when unhappy",
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
			total, _ := usecase.GetTotalPerimeter(&test.canvas, test.id)

			assert.NotEqual(t, test.want, total)

			mockCanvasRepo.AssertExpectations(t)
			mockShapeRepo.AssertExpectations(t)
		})
	}
}

func TestDrawCanvasSuccess(t *testing.T) {
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
			assert.Equal(t, test.want, got)
		})
	}
}

func TestDrawCanvasError(t *testing.T) {
	tests := []struct {
		name   string
		id     string
		canvas models.Canvas
		want   string
	}{
		{
			name: "when unhappy",
			id:   "1",
			canvas: models.Canvas{
				Id:     1,
				Name:   "test",
				Width:  100,
				Height: 100,
				Color:  "#ffffff",
			},
			want: "t.jpg",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			usecase := CanvasUsecase{}
			got, _ := usecase.DrawCanvas(&test.canvas, test.id)
			assert.NotEqual(t, test.want, got)
		})
	}
}
