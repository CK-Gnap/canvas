package controllers

import (
	"canvas/database"
	"canvas/models"
	"errors"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ShapeRepo struct {
	Db *gorm.DB
}

func NewShapeTable() *ShapeRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Shape{})
	return &ShapeRepo{Db: db}
}

func (repository *ShapeRepo) CreateShape(c *gin.Context) {
	canvasID, _ := c.Params.Get("canvas_id")
	canvasIdInt, _ := strconv.ParseInt(canvasID, 10, 64)
	shapeType := c.Param("shape_type")
	if shapeType == string(models.RECTANGLE) {
		repository.createRectangleShape(c, canvasIdInt, shapeType)
	} else if shapeType == string(models.CIRCLE) {
		repository.createCircleShape(c, canvasIdInt, shapeType)
	} else if shapeType == string(models.TRIANGLE) {
		repository.createTriangleShape(c, canvasIdInt, shapeType)
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid shape type"})
		return
	}
}

func (repository *ShapeRepo) createRectangleShape(c *gin.Context, canvasID int64, shapeType string) {
	req := models.Rectangle{}
	if errRequest := c.BindJSON(&req); errRequest != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errRequest.Error()})
		return
	}

	shape := models.Shape{
		CanvasId:  canvasID,
		Type:      models.RECTANGLE,
		X:         req.X,
		Y:         req.Y,
		Width:     req.Width,
		Height:    req.Height,
		Color:     req.Color,
		Area:      req.GetArea(),
		Perimeter: req.GetPerimeter(),
	}

	err := models.CreateShape(&shape, string(canvasID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shape)
}

func (repository *ShapeRepo) createCircleShape(c *gin.Context, canvasID int64, shapeType string) {
	req := models.Circle{}
	if errRequest := c.BindJSON(&req); errRequest != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errRequest.Error()})
		return
	}

	shape := models.Shape{
		CanvasId:  canvasID,
		Type:      models.CIRCLE,
		X:         req.X,
		Y:         req.Y,
		Radius:    req.Radius,
		Color:     req.Color,
		Area:      req.GetArea(),
		Perimeter: req.GetPerimeter(),
	}

	err := models.CreateShape(&shape, string(canvasID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shape)
}

func (repository *ShapeRepo) createTriangleShape(c *gin.Context, canvasID int64, shapeType string) {
	req := models.Triangle{}
	if errRequest := c.BindJSON(&req); errRequest != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errRequest.Error()})
		return
	}

	req.GetSides()
	errTriangle := req.CheckIsTriangle()
	if errTriangle != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errTriangle.Error()})
		return
	}

	shape := models.Shape{
		CanvasId:  canvasID,
		Type:      models.TRIANGLE,
		X:         req.X,
		Y:         req.Y,
		Width:     req.Width,
		Height:    req.Height,
		SideLeft:  req.SideLeft,
		SideRight: req.SideRight,
		SideBase:  req.SideBase,
		Color:     req.Color,
		Area:      req.GetArea(),
		Perimeter: req.GetPerimeter(),
	}

	err := models.CreateShape(&shape, string(canvasID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shape)
}

func (repository *ShapeRepo) GetShape(c *gin.Context) {

	var shape []models.Shape
	canvasID, _ := c.Params.Get("canvas_id")

	err := models.GetShapes(&shape, canvasID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, shape)
}

func (repository *ShapeRepo) GetShapeById(c *gin.Context) {
	var shape models.Shape
	id, _ := c.Params.Get("shape_id")
	_, _, err := models.GetShape(&shape, id)
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

func (repository *ShapeRepo) UpdateShape(c *gin.Context) {
	shape := models.Shape{}
	id, _ := c.Params.Get("shape_id")

	canvasID, shapeType, err := models.GetShape(&shape, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if shapeType == string(models.RECTANGLE) {
		repository.updateRectangleShape(c, canvasID, id)
	} else if shapeType == string(models.CIRCLE) {
		repository.updateCircleShape(c, canvasID, id)
	} else if shapeType == string(models.TRIANGLE) {
		repository.updateTriangleShape(c, canvasID, id)
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid shape type"})
		return
	}
}

func (repository *ShapeRepo) updateRectangleShape(c *gin.Context, canvasID int64, id string) {
	req := models.Rectangle{}
	IdInt, _ := strconv.ParseInt(id, 10, 64)
	if errRequest := c.BindJSON(&req); errRequest != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errRequest.Error()})
		return
	}

	shape := models.Shape{
		Id:        IdInt,
		CanvasId:  canvasID,
		Type:      models.RECTANGLE,
		X:         req.X,
		Y:         req.Y,
		Width:     req.Width,
		Height:    req.Height,
		Color:     req.Color,
		Area:      req.GetArea(),
		Perimeter: req.GetPerimeter(),
	}

	err := models.UpdateShape(&shape)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shape)
}

func (repository *ShapeRepo) updateCircleShape(c *gin.Context, canvasID int64, id string) {
	IdInt, _ := strconv.ParseInt(id, 10, 64)
	req := models.Circle{}
	if errRequest := c.BindJSON(&req); errRequest != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errRequest.Error()})
		return
	}

	shape := models.Shape{
		Id:        IdInt,
		CanvasId:  canvasID,
		Type:      models.CIRCLE,
		X:         req.X,
		Y:         req.Y,
		Radius:    req.Radius,
		Color:     req.Color,
		Area:      req.GetArea(),
		Perimeter: req.GetPerimeter(),
	}

	err := models.UpdateShape(&shape)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shape)
}

func (repository *ShapeRepo) updateTriangleShape(c *gin.Context, canvasID int64, id string) {
	IdInt, _ := strconv.ParseInt(id, 10, 64)
	req := models.Triangle{}
	if errRequest := c.BindJSON(&req); errRequest != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errRequest.Error()})
		return
	}

	req.GetSides()
	errTriangle := req.CheckIsTriangle()
	if errTriangle != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errTriangle.Error()})
		return
	}

	shape := models.Shape{
		Id:        IdInt,
		CanvasId:  canvasID,
		Type:      models.TRIANGLE,
		X:         req.X,
		Y:         req.Y,
		Width:     req.Width,
		Height:    req.Height,
		SideLeft:  req.SideLeft,
		SideRight: req.SideRight,
		SideBase:  req.SideBase,
		Color:     req.Color,
		Area:      req.GetArea(),
		Perimeter: req.GetPerimeter(),
	}

	err := models.UpdateShape(&shape)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shape)
}

func (repository *ShapeRepo) DeleteShape(c *gin.Context) {
	var shape models.Shape
	id, _ := c.Params.Get("shape_id")
	err := models.DeleteShape(&shape, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Shape deleted successfully"})
}

func (repository *ShapeRepo) GetImage(c *gin.Context) {
	var shape []models.Shape
	canvasID, _ := c.Params.Get("canvas_id")
	err := models.GetShapes(&shape, canvasID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	image, err := models.CreateImage(shape[len(shape)-1].Canvas.Name, shape)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fileTmp, errByOpenFile := os.Open(image)
	if errByOpenFile != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errByOpenFile})
		return
	}

	defer fileTmp.Close()

	fileName := path.Base(image)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Disposition", "inline;filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")
	c.File(image)
}
