package repositories_interfaces

import "canvas/models"

type ShapeRepoInterface interface {
	CreateShape(shape *models.Shape, canvasID string) error
	GetShapes(shapes *[]models.Shape, canvasID string) (*[]models.Shape, error)
	GetShape(shape *models.Shape, id string) error
	UpdateShape(shape *models.Shape, id string) (*models.Shape, error)
	DeleteShape(shape *models.Shape, id string) error
}
