package usecases

import (
	"canvas/models"
	models_interfaces "canvas/models/Interfaces"
	repositories_interfaces "canvas/repositories/Interfaces"
	usecases "canvas/usecases/Interfaces"
	"errors"
	"fmt"
	"image"

	"github.com/fogleman/gg"
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
		return nil, errors.New("Error creating canvas")
	}

	return Canvas, nil
}

func (usecase *CanvasUsecase) GetCanvases() ([]models.Canvas, error) {
	var canvases []models.Canvas

	handleCanvasErr := usecase.canvasRepo.GetCanvases(&canvases)
	if handleCanvasErr != nil {
		return nil, errors.New("Error getting canvases")
	}

	for index, canvas := range canvases {
		canvases[index].Shapes = *usecase.getShapes(&canvas)
	}

	return canvases, nil
}

func (usecase *CanvasUsecase) GetCanvas(Canvas *models.Canvas, id string) (*models.Canvas, error) {
	handleCanvasErr := usecase.canvasRepo.GetCanvas(Canvas, id)
	if handleCanvasErr != nil {
		return nil, errors.New("Error getting canvas")
	}

	Canvas.Shapes = *usecase.getShapes(Canvas)

	return Canvas, nil
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

func (usecase *CanvasUsecase) UpdateCanvas(Canvas *models.Canvas, id string) (*models.Canvas, error) {
	var checkCanvas models.Canvas

	handleGetCanvasErr := usecase.canvasRepo.GetCanvas(&checkCanvas, id)
	if handleGetCanvasErr != nil {
		return nil, errors.New("Error getting canvas")
	}

	Canvas.Id = checkCanvas.Id
	handleUpdateCanvasErr := usecase.canvasRepo.UpdateCanvas(Canvas, id)
	if handleUpdateCanvasErr != nil {
		return nil, errors.New("Error updating canvas")
	}

	return Canvas, nil
}

func (usecase *CanvasUsecase) DeleteCanvas(Canvas *models.Canvas, id string) error {
	var checkCanvas models.Canvas

	handleGetCanvasErr := usecase.canvasRepo.GetCanvas(&checkCanvas, id)
	if handleGetCanvasErr != nil {
		return errors.New("Error getting canvas")
	}

	handleCanvasErr := usecase.canvasRepo.DeleteCanvas(Canvas, id)
	if handleCanvasErr != nil {
		return errors.New("Error deleting canvas")
	}

	return handleCanvasErr
}

func (usecase *CanvasUsecase) GetTotalArea(Canvas *models.Canvas, id string) (float64, error) {
	var totalArea float64

	canvas, err := usecase.GetCanvas(Canvas, id)
	if err != nil {
		return 0, errors.New("Error getting canvas")
	}

	for _, shape := range canvas.Shapes {
		totalArea += shape.GetArea()
	}

	return totalArea, nil
}

func (usecase *CanvasUsecase) GetTotalPerimeter(Canvas *models.Canvas, id string) (float64, error) {
	var totalPerimeter float64

	canvas, err := usecase.GetCanvas(Canvas, id)
	if err != nil {
		return 0, err
	}

	for _, shape := range canvas.Shapes {
		totalPerimeter += shape.GetPerimeter()
	}

	return totalPerimeter, nil
}

func (usecase *CanvasUsecase) DrawCanvas(Canvas *models.Canvas, id string) (string, error) {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{int(Canvas.Width), int(Canvas.Height)}
	canvas := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	dc := gg.NewContextForRGBA(canvas)
	dc.SetHexColor(Canvas.Color)
	dc.Clear()

	for _, shape := range Canvas.Shapes {
		dc.Push()
		switch shape.GetType() {
		case string(models.CIRCLE):
			circle := shape.(*models.Circle)
			drawCircle(dc, circle.X, circle.Y, circle.Radius, circle.Color)
		case string(models.RECTANGLE):
			rectangle := shape.(*models.Rectangle)
			drawRectangle(dc, rectangle.X, rectangle.Y, rectangle.Width, rectangle.Height, rectangle.Color)
		case string(models.TRIANGLE):
			triangle := shape.(*models.Triangle)
			drawTriangle(dc, triangle.X, triangle.Y, triangle.Width, triangle.Height, triangle.Color)
		default:
			return "", errors.New("Invalid shape type")
		}
		dc.Pop()
	}

	canvasName := Canvas.Name + ".jpg"

	dc.SaveJPG(canvasName, 100)

	return canvasName, nil
}

func drawCircle(dc *gg.Context, x, y, radius float64, color string) {
	dc.DrawCircle(x, y, radius)
	// dc.SetHexColor("ffffff")
	// dc.SetFillRule(gg.FillRuleEvenOdd)
	// dc.FillPreserve()
	dc.SetHexColor(color)
	dc.SetLineWidth(1)
	dc.Stroke()
}

func drawRectangle(dc *gg.Context, x, y, width, height float64, color string) {
	dc.DrawRectangle(x, y, width, height)
	// dc.SetHexColor("ffffff")
	// dc.SetFillRule(gg.FillRuleEvenOdd)
	// dc.FillPreserve()
	dc.SetHexColor(color)
	dc.SetLineWidth(1)
	dc.Stroke()
}

func drawTriangle(dc *gg.Context, x, y, width, height float64, color string) {
	dc.DrawLine(x, y, x+width, y)
	dc.DrawLine(x+width, y, x+width/2, y+height)
	dc.DrawLine(x, y, x+width/2, y+height)
	// dc.SetHexColor("ffffff")
	// dc.SetFillRule(gg.FillRuleEvenOdd)
	// dc.FillPreserve()
	dc.SetHexColor(color)
	dc.SetLineWidth(1)
	dc.Stroke()
}
