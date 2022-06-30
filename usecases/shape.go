package usecases

import (
	models "canvas/models"
	models_interfaces "canvas/models/Interfaces"
	repositories_interfaces "canvas/repositories/Interfaces"
	usecases_interfaces "canvas/usecases/Interfaces"
	"errors"
	"strconv"
)

type ShapeUsecase struct {
	canvasRepo repositories_interfaces.CanvasRepoInterface
	shapeRepo  repositories_interfaces.ShapeRepoInterface
}

func NewShapeUsecase(canvasRepoInterface repositories_interfaces.CanvasRepoInterface, shapeRepoInterface repositories_interfaces.ShapeRepoInterface) usecases_interfaces.ShapeUsecaseInterface {
	return &ShapeUsecase{
		canvasRepo: canvasRepoInterface,
		shapeRepo:  shapeRepoInterface,
	}
}

func (usecase *ShapeUsecase) CreateRectangleShape(Rectangle *models.Rectangle, canvasID string) (models_interfaces.ShapeInterface, error) {
	var canvas models.Canvas
	canvasIdInt, _ := strconv.ParseInt(canvasID, 10, 64)

	handleGetCanvasErr := usecase.canvasRepo.GetCanvas(&canvas, canvasID)
	if handleGetCanvasErr != nil {
		return nil, errors.New("Error getting canvas")
	}

	shape := models.Shape{
		Type:     models.RECTANGLE,
		CanvasId: canvasIdInt,
		X:        Rectangle.X,
		Y:        Rectangle.Y,
		Width:    Rectangle.Width,
		Height:   Rectangle.Height,
		Color:    Rectangle.Color,
	}

	handleShapeErr := usecase.shapeRepo.CreateShape(&shape, canvasID)
	if handleShapeErr != nil {
		return nil, errors.New("Error creating shape")
	}

	return ConvertTypeOfShape(&shape), nil
}

func (usecase *ShapeUsecase) CreateCircleShape(Circle *models.Circle, canvasID string) (models_interfaces.ShapeInterface, error) {
	var canvas models.Canvas
	canvasIdInt, _ := strconv.ParseInt(canvasID, 10, 64)

	handleGetCanvasErr := usecase.canvasRepo.GetCanvas(&canvas, canvasID)
	if handleGetCanvasErr != nil {
		return nil, handleGetCanvasErr
	}

	shape := models.Shape{
		Type:     models.CIRCLE,
		CanvasId: canvasIdInt,
		X:        Circle.X,
		Y:        Circle.Y,
		Radius:   Circle.Radius,
		Color:    Circle.Color,
	}

	handleShapeErr := usecase.shapeRepo.CreateShape(&shape, canvasID)
	if handleShapeErr != nil {
		return nil, handleShapeErr
	}

	return ConvertTypeOfShape(&shape), nil
}

func (usecase *ShapeUsecase) CreateTriangleShape(Triangle *models.Triangle, canvasID string) (models_interfaces.ShapeInterface, error) {
	var canvas models.Canvas
	canvasIdInt, _ := strconv.ParseInt(canvasID, 10, 64)

	handleGetCanvasErr := usecase.canvasRepo.GetCanvas(&canvas, canvasID)
	if handleGetCanvasErr != nil {
		return nil, handleGetCanvasErr
	}

	shape := models.Shape{
		Type:     models.TRIANGLE,
		CanvasId: canvasIdInt,
		X:        Triangle.X,
		Y:        Triangle.Y,
		Width:    Triangle.Width,
		Height:   Triangle.Height,
		Color:    Triangle.Color,
	}

	handleShapeErr := usecase.shapeRepo.CreateShape(&shape, canvasID)
	if handleShapeErr != nil {
		return nil, handleShapeErr
	}

	return ConvertTypeOfShape(&shape), nil
}

func (usecase *ShapeUsecase) GetShapes(canvasID string) ([]models_interfaces.ShapeInterface, error) {
	var canvas models.Canvas
	var shapesModel []models.Shape

	handleGetCanvasErr := usecase.canvasRepo.GetCanvas(&canvas, canvasID)
	if handleGetCanvasErr != nil {
		return nil, errors.New("Error getting canvas")
	}

	shapes, handleShapeErr := usecase.shapeRepo.GetShapes(&shapesModel, canvasID)
	if handleShapeErr != nil {
		return nil, errors.New("Error getting shapes")
	}
	return ConvertTypeOfShapes(shapes), nil
}

func (usecase *ShapeUsecase) GetShape(Shape *models.Shape, id string) (models_interfaces.ShapeInterface, error) {
	handleShapeErr := usecase.shapeRepo.GetShape(Shape, id)
	if handleShapeErr != nil {
		return nil, errors.New("Error getting shape")
	}

	return ConvertTypeOfShape(Shape), nil
}

func (usecase *ShapeUsecase) UpdateShape(Shape *models.Shape, id string) (models_interfaces.ShapeInterface, error) {
	var checkShape models.Shape

	handleGetShapeErr := usecase.shapeRepo.GetShape(&checkShape, id)
	if handleGetShapeErr != nil {
		return nil, errors.New("Error getting shape")
	}

	Shape.Id = checkShape.Id
	Shape.CanvasId = checkShape.CanvasId
	Shape.Type = checkShape.Type
	update, handleUpdateShapeErr := usecase.shapeRepo.UpdateShape(Shape, id)
	if handleUpdateShapeErr != nil {
		return nil, errors.New("Error updating shape")
	}

	return ConvertTypeOfShape(update), nil
}

func (usecase *ShapeUsecase) DeleteShape(Shape *models.Shape, id string) error {
	var checkShape models.Shape

	handleGetShapeErr := usecase.shapeRepo.GetShape(&checkShape, id)
	if handleGetShapeErr != nil {
		return errors.New("Error getting shape")
	}

	handleShapeErr := usecase.shapeRepo.DeleteShape(Shape, id)
	if handleShapeErr != nil {
		return errors.New("Error deleting shape")
	}
	return handleShapeErr
}

func ConvertTypeOfShapes(Shapes *[]models.Shape) []models_interfaces.ShapeInterface {
	shapeInterfaces := []models_interfaces.ShapeInterface{}

	for _, Shape := range *Shapes {
		shapeInterfaces = append(shapeInterfaces, ConvertTypeOfShape(&Shape))
	}

	return shapeInterfaces
}

func ConvertTypeOfShape(Shape *models.Shape) models_interfaces.ShapeInterface {
	switch Shape.Type {
	case models.RECTANGLE:
		return models.ConvertToRectangle(Shape)
	case models.CIRCLE:
		return models.ConvertToCircle(Shape)
	case models.TRIANGLE:
		return models.ConvertToTriangle(Shape)
	default:
		return nil
	}
}
