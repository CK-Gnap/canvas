package deliveries

import (
	"canvas/models"
	usecases "canvas/usecases/Interfaces"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CanvasHandler struct {
	CanvasUsecase usecases.CanvasUsecaseInterface
}

func (handler *CanvasHandler) CreateCanvas(c *gin.Context) {
	var req models.Canvas

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	canvas, err := handler.CanvasUsecase.CreateCanvas(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, canvas)
}

func (handler *CanvasHandler) GetCanvases(c *gin.Context) {
	canvas, err := handler.CanvasUsecase.GetCanvases()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, canvas)
}

func (handler *CanvasHandler) GetCanvas(c *gin.Context) {
	id, _ := c.Params.Get("canvas_id")
	var req models.Canvas

	canvas, err := handler.CanvasUsecase.GetCanvas(&req, id)
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

func (handler *CanvasHandler) UpdateCanvas(c *gin.Context) {
	var req models.Canvas
	id, _ := c.Params.Get("canvas_id")

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	canvas, err := handler.CanvasUsecase.UpdateCanvas(&req, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, canvas)
}

func (handler *CanvasHandler) DeleteCanvas(c *gin.Context) {
	var req models.Canvas
	id, _ := c.Params.Get("canvas_id")

	err := handler.CanvasUsecase.DeleteCanvas(&req, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Canvas deleted successfully"})
}

// func (handler *CanvasHandler) GetImage(c *gin.Context) {
// 	var shape []models.ShapeInterface
// 	image, err := models.CreateImage("image", shape)

// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	fileTmp, errByOpenFile := os.Open(image)
// 	if errByOpenFile != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errByOpenFile})
// 		return
// 	}

// 	defer fileTmp.Close()

// 	fileName := path.Base(image)
// 	c.Header("Content-Type", "application/octet-stream")
// 	c.Header("Content-Disposition", "attachment; filename="+fileName)
// 	c.Header("Content-Disposition", "inline;filename="+fileName)
// 	c.Header("Content-Transfer-Encoding", "binary")
// 	c.Header("Cache-Control", "no-cache")
// 	c.File(image)
// }
