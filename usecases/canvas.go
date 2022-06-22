package usecases

import (
	"canvas/models"
	models_interfaces "canvas/models/Interfaces"
	repositories_interfaces "canvas/repositories/Interfaces"
	usecases "canvas/usecases/Interfaces"
	"fmt"
)

type CanvasUsecase struct {
	canvasRepo repositories_interfaces.CanvasRepoInterface
	shapeRepo  repositories_interfaces.ShapeRepoInterface
}

func NewCanvasUsecase(canvasRepoInterface repositories_interfaces.CanvasRepoInterface, shapeRepoInterface repositories_interfaces.ShapeRepoInterface) usecases.CanvasUsecaseInterface {
	return &CanvasUsecase{
		canvasRepo: canvasRepoInterface,
		shapeRepo:  shapeRepoInterface,
	}
}

func (usecase *CanvasUsecase) CreateCanvas(Canvas *models.Canvas) (*models.Canvas, error) {
	handleCanvasErr := usecase.canvasRepo.CreateCanvas(Canvas)
	if handleCanvasErr != nil {
		return nil, handleCanvasErr
	}

	return Canvas, nil
}

func (usecase *CanvasUsecase) GetCanvases() ([]models.Canvas, error) {
	var canvases []models.Canvas

	handleCanvasErr := usecase.canvasRepo.GetCanvases(&canvases)
	if handleCanvasErr != nil {
		return nil, handleCanvasErr
	}

	for index, canvas := range canvases {
		canvases[index].Shapes = *usecase.getShapes(&canvas)
	}

	return canvases, nil
}

func (usecase *CanvasUsecase) GetCanvas(Canvas *models.Canvas, id string) (*models.Canvas, error) {
	handleCanvasErr := usecase.canvasRepo.GetCanvas(Canvas, id)
	if handleCanvasErr != nil {
		return nil, handleCanvasErr
	}

	Canvas.Shapes = *usecase.getShapes(Canvas)

	return Canvas, nil
}

func (usecase *CanvasUsecase) UpdateCanvas(Canvas *models.Canvas, id string) (*models.Canvas, error) {
	var checkCanvas models.Canvas

	handleGetCanvasErr := usecase.canvasRepo.GetCanvas(&checkCanvas, id)
	if handleGetCanvasErr != nil {
		return nil, handleGetCanvasErr
	}

	handleUpdateCanvasErr := usecase.canvasRepo.UpdateCanvas(Canvas, id)
	if handleUpdateCanvasErr != nil {
		return nil, handleUpdateCanvasErr
	}

	return Canvas, nil
}

func (usecase *CanvasUsecase) DeleteCanvas(Canvas *models.Canvas, id string) error {
	var checkCanvas models.Canvas

	handleGetCanvasErr := usecase.canvasRepo.GetCanvas(&checkCanvas, id)
	if handleGetCanvasErr != nil {
		return handleGetCanvasErr
	}

	handleCanvasErr := usecase.canvasRepo.DeleteCanvas(Canvas, id)

	return handleCanvasErr
}

func (usecase *CanvasUsecase) getShapes(Canvas *models.Canvas) *[]models_interfaces.ShapeInterface {
	canvasShapes := []models_interfaces.ShapeInterface{}
	var shapes []models.Shape

	getShapes, _ := usecase.shapeRepo.GetShapes(&shapes, fmt.Sprintf("%v", Canvas.Id))
	for _, shape := range *getShapes {
		switch shape.Type {
		case models.RECTANGLE:
			canvasShapes = append(canvasShapes, models.ConvertToRectangle(&shape))
		case models.CIRCLE:
			canvasShapes = append(canvasShapes, models.ConvertToCircle(&shape))
		case models.TRIANGLE:
			canvasShapes = append(canvasShapes, models.ConvertToTriangle(&shape))
		default:
			return nil
		}
	}

	return &canvasShapes
}
