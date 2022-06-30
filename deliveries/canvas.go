package deliveries

import (
	"canvas/constants"
	models "canvas/models"
	usecases "canvas/usecases/Interfaces"
	"errors"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CanvasHandler struct {
	CanvasUsecase usecases.CanvasUsecaseInterface
}

func (handler *CanvasHandler) CreateCanvas(c *gin.Context) {
	var req models.CanvasRequestCreate

	if err := c.BindJSON(&req); err != nil {
		responseError(c, http.StatusBadRequest, err)
		return
	}

	newReq := models.Canvas{
		Name:   req.Name,
		Width:  req.Width,
		Height: req.Height,
		Color:  req.Color,
	}

	canvas, err := handler.CanvasUsecase.CreateCanvas(&newReq)
	if err != nil {
		responseError(c, http.StatusInternalServerError, err)
		return
	}

	responseSuccess(c, http.StatusOK, constants.SuccessCreateCavas, canvas)
}

func (handler *CanvasHandler) GetCanvases(c *gin.Context) {
	canvas, err := handler.CanvasUsecase.GetCanvases()
	if err != nil {
		responseError(c, http.StatusInternalServerError, err)
		return
	}

	responseSuccess(c, http.StatusOK, constants.SuccessGetCanvases, canvas)
}

func (handler *CanvasHandler) GetCanvas(c *gin.Context) {
	id, _ := c.Params.Get("canvas_id")
	var req models.Canvas

	canvas, err := handler.CanvasUsecase.GetCanvas(&req, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responseError(c, http.StatusNotFound, err)
			return
		}

		responseError(c, http.StatusInternalServerError, err)
		return
	}

	responseSuccess(c, http.StatusOK, constants.SuccessGetCanvas, canvas)
}

func (handler *CanvasHandler) UpdateCanvas(c *gin.Context) {
	var req models.CanvasRequestUpdare
	id, _ := c.Params.Get("canvas_id")

	if err := c.BindJSON(&req); err != nil {
		responseError(c, http.StatusBadRequest, err)
		return
	}

	newReq := models.Canvas{
		Name:   req.Name,
		Width:  req.Width,
		Height: req.Height,
		Color:  req.Color,
	}

	canvas, err := handler.CanvasUsecase.UpdateCanvas(&newReq, id)
	if err != nil {
		responseError(c, http.StatusInternalServerError, err)
		return
	}

	responseSuccess(c, http.StatusOK, constants.SuccessUpdateCanvas, canvas)
}

func (handler *CanvasHandler) DeleteCanvas(c *gin.Context) {
	var req models.Canvas
	id, _ := c.Params.Get("canvas_id")

	err := handler.CanvasUsecase.DeleteCanvas(&req, id)
	if err != nil {
		responseError(c, http.StatusInternalServerError, err)
		return
	}

	responseSuccess(c, http.StatusOK, constants.SuccessDeleteCanvas, nil)
}

func (handler *CanvasHandler) GetTotalArea(c *gin.Context) {
	var req models.Canvas
	id, _ := c.Params.Get("canvas_id")

	totalArea, err := handler.CanvasUsecase.GetTotalArea(&req, id)
	if err != nil {
		responseError(c, http.StatusInternalServerError, err)
		return
	}

	responseSuccess(c, http.StatusOK, constants.SuccessTotalArea, totalArea)
}

func (handler *CanvasHandler) GetTotalPerimeter(c *gin.Context) {
	var req models.Canvas
	id, _ := c.Params.Get("canvas_id")

	totalPerimeter, err := handler.CanvasUsecase.GetTotalPerimeter(&req, id)
	if err != nil {
		responseError(c, http.StatusInternalServerError, err)
		return
	}

	responseSuccess(c, http.StatusOK, constants.SuccessTotalPerimeter, totalPerimeter)
}

func (handler *CanvasHandler) DrawCanvas(c *gin.Context) {
	id, _ := c.Params.Get("canvas_id")
	var req models.Canvas

	canvas, err := handler.CanvasUsecase.GetCanvas(&req, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responseError(c, http.StatusNotFound, err)
			return
		}
		responseError(c, http.StatusInternalServerError, err)
		return
	}

	image, errDrawCanvas := handler.CanvasUsecase.DrawCanvas(canvas, id)
	if errDrawCanvas != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errDrawCanvas})
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
