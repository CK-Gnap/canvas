package deliveries

import (
	"canvas/constants"
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
		responseError(c, http.StatusBadRequest, errRequest)
		return
	}

	newReq := models.Rectangle(req)

	shape, err := handler.ShapeUsecase.CreateRectangleShape(&newReq, canvasID)
	if err != nil {
		responseError(c, http.StatusInternalServerError, err)
		return
	}

	responseSuccess(c, http.StatusOK, constants.SuccessCreateShape, shape)
}

func (handler *ShapeHandler) CreateCircleShape(c *gin.Context) {
	canvasID, _ := c.Params.Get("canvas_id")
	var req models.CircleRequestCreate

	if errRequest := c.BindJSON(&req); errRequest != nil {
		responseError(c, http.StatusBadRequest, errRequest)
		return
	}

	newReq := models.Circle(req)

	shape, err := handler.ShapeUsecase.CreateCircleShape(&newReq, canvasID)
	if err != nil {
		responseError(c, http.StatusInternalServerError, err)
		return
	}

	responseSuccess(c, http.StatusOK, constants.SuccessCreateShape, shape)

}

func (handler *ShapeHandler) CreateTriangleShape(c *gin.Context) {
	canvasID, _ := c.Params.Get("canvas_id")
	var req models.TriangleRequestCreate

	if errRequest := c.BindJSON(&req); errRequest != nil {
		responseError(c, http.StatusBadRequest, errRequest)
		return
	}

	newReq := models.Triangle(req)

	shape, err := handler.ShapeUsecase.CreateTriangleShape(&newReq, canvasID)
	if err != nil {
		responseError(c, http.StatusInternalServerError, err)
		return
	}

	responseSuccess(c, http.StatusOK, constants.SuccessCreateShape, shape)
}

func (handler *ShapeHandler) GetShapes(c *gin.Context) {
	canvasID, _ := c.Params.Get("canvas_id")

	shape, err := handler.ShapeUsecase.GetShapes(canvasID)
	if err != nil {
		responseError(c, http.StatusInternalServerError, err)
		return
	}

	responseSuccess(c, http.StatusOK, constants.SuccessGetShapes, shape)
}

func (handler *ShapeHandler) GetShape(c *gin.Context) {
	id, _ := c.Params.Get("shape_id")
	var req models.Shape

	shape, err := handler.ShapeUsecase.GetShape(&req, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responseError(c, http.StatusNotFound, err)
			return
		}
		responseError(c, http.StatusInternalServerError, err)
		return
	}
	responseSuccess(c, http.StatusOK, constants.SuccessGetShape, shape)

}

func (handler *ShapeHandler) UpdateShape(c *gin.Context) {
	id, _ := c.Params.Get("shape_id")
	var req models.ShapeRequestUpdate
	if err := c.BindJSON(&req); err != nil {
		responseError(c, http.StatusBadRequest, err)
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
		responseError(c, http.StatusInternalServerError, errUpdareShape)
		return
	}

	responseSuccess(c, http.StatusOK, constants.SuccessUpdateShape, shape)

}

func (handler *ShapeHandler) DeleteShape(c *gin.Context) {
	var req models.Shape
	id, _ := c.Params.Get("shape_id")
	err := handler.ShapeUsecase.DeleteShape(&req, id)
	if err != nil {
		responseError(c, http.StatusInternalServerError, err)
		return
	}
	responseSuccess(c, http.StatusOK, constants.SuccessDeleteShape, nil)
}
