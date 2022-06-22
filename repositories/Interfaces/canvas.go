package repositories_interfaces

import "canvas/models"

type CanvasRepoInterface interface {
	CreateCanvas(canvas *models.Canvas) error
	GetCanvases(canvases *[]models.Canvas) error
	GetCanvas(canvas *models.Canvas, id string) error
	UpdateCanvas(canvas *models.Canvas, id string) error
	DeleteCanvas(canvas *models.Canvas, id string) error
}
