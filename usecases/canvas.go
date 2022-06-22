package usecases

import (
	"canvas/models"
	repositories_interfaces "canvas/repositories/Interfaces"
	usecases "canvas/usecases/Interfaces"
)

type CanvasUsecase struct {
	repo repositories_interfaces.CanvasRepoInterface
}

func NewCanvasUsecase(canvasRepoInterface repositories_interfaces.CanvasRepoInterface) usecases.CanvasUsecaseInterface {
	return &CanvasUsecase{
		repo: canvasRepoInterface,
	}
}

func (usecase *CanvasUsecase) CreateCanvas(Canvas *models.Canvas) (*models.Canvas, error) {
	handleCanvasErr := usecase.repo.CreateCanvas(Canvas)
	if handleCanvasErr != nil {
		return nil, handleCanvasErr
	}

	return Canvas, nil
}

func (usecase *CanvasUsecase) GetCanvases() ([]models.Canvas, error) {
	var canvases []models.Canvas

	handleCanvasErr := usecase.repo.GetCanvases(&canvases)
	if handleCanvasErr != nil {
		return nil, handleCanvasErr
	}

	return canvases, nil
}

func (usecase *CanvasUsecase) GetCanvas(Canvas *models.Canvas, id string) (*models.Canvas, error) {
	handleCanvasErr := usecase.repo.GetCanvas(Canvas, id)
	if handleCanvasErr != nil {
		return nil, handleCanvasErr
	}

	return Canvas, nil
}

func (usecase *CanvasUsecase) UpdateCanvas(Canvas *models.Canvas, id string) (*models.Canvas, error) {
	var checkCanvas models.Canvas

	handleGetCanvasErr := usecase.repo.GetCanvas(&checkCanvas, id)
	if handleGetCanvasErr != nil {
		return nil, handleGetCanvasErr
	}

	handleUpdateCanvasErr := usecase.repo.UpdateCanvas(Canvas, id)
	if handleUpdateCanvasErr != nil {
		return nil, handleUpdateCanvasErr
	}

	return Canvas, nil
}

func (usecase *CanvasUsecase) DeleteCanvas(Canvas *models.Canvas, id string) error {
	var checkCanvas models.Canvas

	handleGetCanvasErr := usecase.repo.GetCanvas(&checkCanvas, id)
	if handleGetCanvasErr != nil {
		return handleGetCanvasErr
	}

	handleCanvasErr := usecase.repo.DeleteCanvas(Canvas, id)

	return handleCanvasErr
}
