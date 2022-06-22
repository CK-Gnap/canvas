package routes

import (
	"canvas/database"
	"canvas/deliveries"
	"canvas/repositories"
	"canvas/usecases"
	usecases_interfaces "canvas/usecases/Interfaces"

	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	//* repositories
	canvasRepo := repositories.NewCanvasRepo(database.InitDb())
	shapeRepo := repositories.NewShapeRepo(database.InitDb())

	//* usecases
	canvasUsecase := usecases.NewCanvasUsecase(canvasRepo, shapeRepo)
	shapeUsecase := usecases.NewShapeUsecase(canvasRepo, shapeRepo)

	//* deliveries
	canvasDelivery(r, canvasUsecase)
	shapeDelivery(r, shapeUsecase)

	return r
}

func canvasDelivery(r *gin.Engine, usecase usecases_interfaces.CanvasUsecaseInterface) {

	handler := &deliveries.CanvasHandler{
		CanvasUsecase: usecase,
	}

	r.POST("/canvas", handler.CreateCanvas)
	r.GET("/canvas", handler.GetCanvases)
	r.GET("/canvas/:canvas_id", handler.GetCanvas)
	r.PUT("/canvas/:canvas_id", handler.UpdateCanvas)
	r.DELETE("/canvas/:canvas_id", handler.DeleteCanvas)
}

func shapeDelivery(r *gin.Engine, usecase usecases_interfaces.ShapeUsecaseInterface) {

	handler := &deliveries.ShapeHandler{
		ShapeUsecase: usecase,
	}

	r.POST("/canvas/:canvas_id/rectangle", handler.CreateRectangleShape)
	r.POST("/canvas/:canvas_id/circle", handler.CreateCircleShape)
	r.POST("/canvas/:canvas_id/triangle", handler.CreateTriangleShape)
	r.GET("/canvas/:canvas_id/shapes", handler.GetShapes)
	r.GET("/canvas/shapes/:shape_id", handler.GetShape)
	r.PUT("/canvas/shapes/:shape_id", handler.UpdateShape)
	r.DELETE("/canvas/shapes/:shape_id", handler.DeleteShape)
	// r.GET("/canvas/:canvas_id/image", handler.GetImage)
}
