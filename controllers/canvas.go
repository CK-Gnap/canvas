package controllers

import (
	"canvas/database"
	"canvas/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CanvasRepo struct {
	Db *gorm.DB
}

func NewCanvasTable() *CanvasRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Canvas{})
	return &CanvasRepo{Db: db}
}

func (repository *CanvasRepo) CreateCanvas(c *gin.Context) {
	var canvas models.Canvas
	c.BindJSON(&canvas)
	err := models.CreateCanvas(repository.Db, &canvas)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, canvas)
}

func (repository *CanvasRepo) GetCanvas(c *gin.Context) {
	var canvas []models.Canvas
	err := models.GetCanvases(repository.Db, &canvas)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, canvas)
}

func (repository *CanvasRepo) GetCanvasById(c *gin.Context) {
	id, _ := c.Params.Get("canvas_id")
	var canvas models.Canvas
	err := models.GetCanvas(repository.Db, &canvas, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, canvas)
}

func (repository *CanvasRepo) UpdateCanvas(c *gin.Context) {
	var canvas models.Canvas
	id, _ := c.Params.Get("canvas_id")
	err := models.GetCanvas(repository.Db, &canvas, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&canvas)
	err = models.UpdateCanvas(repository.Db, &canvas)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, canvas)
}

func (repository *CanvasRepo) DeleteCanvas(c *gin.Context) {
	var canvas models.Canvas
	id, _ := c.Params.Get("canvas_id")
	err := models.DeleteCanvas(repository.Db, &canvas, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Canvas deleted successfully"})
}
