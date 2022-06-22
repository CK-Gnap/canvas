package deliveries

import (
	"canvas/models"

	usecases "canvas/usecases/Interfaces"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ShapeHandler struct {
	ShapeUsecase usecases.ShapeUsecaseInterface
}

func (handler *ShapeHandler) CreateRectangleShape(c *gin.Context) {
	canvasID, _ := c.Params.Get("canvas_id")
	var req models.RectangleRequestCreate

	if errRequest := c.BindJSON(&req); errRequest != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errRequest.Error()})
		return
	}

	newReq := models.Rectangle(req)

	shape, err := handler.ShapeUsecase.CreateRectangleShape(&newReq, canvasID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shape)
}

func (handler *ShapeHandler) CreateCircleShape(c *gin.Context) {
	canvasID, _ := c.Params.Get("canvas_id")
	var req models.CircleRequestCreate

	if errRequest := c.BindJSON(&req); errRequest != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errRequest.Error()})
		return
	}

	newReq := models.Circle(req)

	shape, err := handler.ShapeUsecase.CreateCircleShape(&newReq, canvasID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shape)
}

func (handler *ShapeHandler) CreateTriangleShape(c *gin.Context) {
	canvasID, _ := c.Params.Get("canvas_id")
	var req models.TriangleRequestCreate

	if errRequest := c.BindJSON(&req); errRequest != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errRequest.Error()})
		return
	}

	newReq := models.Triangle(req)

	shape, err := handler.ShapeUsecase.CreateTriangleShape(&newReq, canvasID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shape)
}

func (handler *ShapeHandler) GetShapes(c *gin.Context) {
	canvasID, _ := c.Params.Get("canvas_id")

	shape, err := handler.ShapeUsecase.GetShapes(canvasID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shape)
}

func (handler *ShapeHandler) GetShape(c *gin.Context) {
	id, _ := c.Params.Get("shape_id")
	var req models.Shape

	shape, err := handler.ShapeUsecase.GetShape(&req, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, shape)
}

func (handler *ShapeHandler) UpdateShape(c *gin.Context) {
	id, _ := c.Params.Get("shape_id")
	var req models.ShapeRequestUpdate
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newReq := models.Shape{
		X:      req.X,
		Y:      req.Y,
		Width:  req.Width,
		Height: req.Height,
		Radius: req.Radius,
		Color:  req.Color,
	}

	shape, errUpdareShape := handler.ShapeUsecase.UpdateShape(&newReq, id)
	if errUpdareShape != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errUpdareShape.Error()})
		return
	}

	c.JSON(http.StatusOK, shape)
}

func (handler *ShapeHandler) DeleteShape(c *gin.Context) {
	var req models.Shape
	id, _ := c.Params.Get("shape_id")
	err := handler.ShapeUsecase.DeleteShape(&req, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Shape deleted successfully"})
}
