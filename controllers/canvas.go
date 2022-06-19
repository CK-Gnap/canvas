package controllers

import (
	"canvas/database"
	"canvas/models"
	"errors"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CanvasRepo struct {
	Db     *gorm.DB
	Canvas models.Canvas
}

func NewCanvasTable() *CanvasRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Canvas{})
	return &CanvasRepo{Db: db}
}

func (repository *CanvasRepo) CreateCanvas(c *gin.Context) {
	var canvas models.Canvas
	if err := c.BindJSON(&canvas); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err := models.CreateCanvas(&canvas)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, canvas)
}

func (repository *CanvasRepo) GetCanvas(c *gin.Context) {
	var canvas []models.Canvas
	err := models.GetCanvas(&canvas)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, canvas)
}

func (repository *CanvasRepo) GetCanvasById(c *gin.Context) {
	id, _ := c.Params.Get("canvas_id")
	var canvas models.Canvas
	err := models.GetCanvasById(&canvas, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, canvas)
}

func (repository *CanvasRepo) UpdateCanvas(c *gin.Context) {
	var canvas models.Canvas
	id, _ := c.Params.Get("canvas_id")
	err := models.GetCanvasById(&canvas, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := c.BindJSON(&canvas); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err = models.UpdateCanvas(&canvas)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, canvas)
}

func (repository *CanvasRepo) DeleteCanvas(c *gin.Context) {
	var canvas models.Canvas
	id, _ := c.Params.Get("canvas_id")
	err := models.DeleteCanvas(&canvas, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Canvas deleted successfully"})
}

func (repository *CanvasRepo) GetImage(c *gin.Context) {
	var shape []models.Shape
	canvasID, _ := c.Params.Get("canvas_id")
	err := models.GetShapes(&shape, canvasID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	image, err := repository.Canvas.CreateImage()

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
