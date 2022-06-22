package usecases

import (
	models "canvas/models"
	models_interfaces "canvas/models/Interfaces"
)

type ShapeUsecaseInterface interface {
	CreateRectangleShape(shape *models.Rectangle, canvasID string) (models_interfaces.ShapeInterface, error)
	CreateCircleShape(shape *models.Circle, canvasID string) (models_interfaces.ShapeInterface, error)
	CreateTriangleShape(shape *models.Triangle, canvasID string) (models_interfaces.ShapeInterface, error)
	GetShapes(canvasID string) ([]models_interfaces.ShapeInterface, error)
	GetShape(shape *models.Shape, id string) (models_interfaces.ShapeInterface, error)
	UpdateShape(shape *models.Shape, id string) (models_interfaces.ShapeInterface, error)
	DeleteShape(shape *models.Shape, id string) error
}
