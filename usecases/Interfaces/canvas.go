package usecases

import "canvas/models"

type CanvasUsecaseInterface interface {
	CreateCanvas(canvas *models.Canvas) (*models.Canvas, error)
	GetCanvases() (canvases []models.Canvas, err error)
	GetCanvas(canvas *models.Canvas, id string) (*models.Canvas, error)
	UpdateCanvas(canvas *models.Canvas, id string) (*models.Canvas, error)
	DeleteCanvas(canvas *models.Canvas, id string) error
}
