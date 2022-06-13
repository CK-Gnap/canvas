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

type ShapeRepo struct {
	Db *gorm.DB
}

func NewShapeTable() *ShapeRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Shape{})
	return &ShapeRepo{Db: db}
}

func (repository *ShapeRepo) CreateShape(c *gin.Context) {
	var shape models.Shape
	canvasID, _ := c.Params.Get("canvas_id")
	c.BindJSON(&shape)

	var err error
	if shape.Type == "rectangle" {
		shapeType := models.Rectangle{}
		err = shapeType.CreateShape(repository.Db, &shape, canvasID)
	} else if shape.Type == "circle" {
		shapeType := models.Circle{}
		err = shapeType.CreateShape(repository.Db, &shape, canvasID)
	} else if shape.Type == "triangle" {
		shapeType := models.Triangle{}
		err = shapeType.CreateShape(repository.Db, &shape, canvasID)
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shape)
}

func (repository *ShapeRepo) GetShape(c *gin.Context) {

	var shape []models.Shape
	canvasID, _ := c.Params.Get("canvas_id")
	err := models.GetShapes(repository.Db, &shape, canvasID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, shape)
}

func (repository *ShapeRepo) GetShapeById(c *gin.Context) {
	var shape models.Shape
	id, _ := c.Params.Get("shape_id")
	err := models.GetShape(repository.Db, &shape, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, shape)
}

func (repository *ShapeRepo) UpdateShape(c *gin.Context) {
	var shape models.Shape
	id, _ := c.Params.Get("shape_id")

	err := models.GetShape(repository.Db, &shape, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&shape)

	var errUpdate error
	if shape.Type == "rectangle" {
		shapeType := models.Rectangle{}
		errUpdate = shapeType.UpdateShape(repository.Db, &shape)
	} else if shape.Type == "circle" {
		shapeType := models.Circle{}
		errUpdate = shapeType.UpdateShape(repository.Db, &shape)
	} else if shape.Type == "triangle" {
		shapeType := models.Triangle{}
		errUpdate = shapeType.UpdateShape(repository.Db, &shape)
	}

	if errUpdate != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errUpdate})
		return
	}
	c.JSON(http.StatusOK, shape)
}

func (repository *ShapeRepo) DeleteShape(c *gin.Context) {
	var shape models.Shape
	id, _ := c.Params.Get("shape_id")
	err := models.DeleteShape(repository.Db, &shape, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Shape deleted successfully"})
}

func (repository *ShapeRepo) GetImage(c *gin.Context) {
	var shape []models.Shape
	canvasID, _ := c.Params.Get("canvas_id")
	err := models.GetShapes(repository.Db, &shape, canvasID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	image, err := models.CreateImage(shape)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
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
