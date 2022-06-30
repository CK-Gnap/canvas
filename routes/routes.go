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

	canvas := r.Group("/canvas")
	canvas.POST("", handler.CreateCanvas)
	canvas.GET("", handler.GetCanvases)
	canvas.GET("/:canvas_id", handler.GetCanvas)
	canvas.PUT("/:canvas_id", handler.UpdateCanvas)
	canvas.DELETE("/:canvas_id", handler.DeleteCanvas)
	canvas.GET("/:canvas_id/totalArea", handler.GetTotalArea)
	canvas.GET("/:canvas_id/totalPerimeter", handler.GetTotalPerimeter)
	canvas.GET("/:canvas_id/image", handler.DrawCanvas)
}

func shapeDelivery(r *gin.Engine, usecase usecases_interfaces.ShapeUsecaseInterface) {

	handler := &deliveries.ShapeHandler{
		ShapeUsecase: usecase,
	}

	canvas := r.Group("/canvas")
	canvas.POST("/:canvas_id/rectangle", handler.CreateRectangleShape)
	canvas.POST("/:canvas_id/circle", handler.CreateCircleShape)
	canvas.POST("/:canvas_id/triangle", handler.CreateTriangleShape)
	canvas.GET("/:canvas_id/shapes", handler.GetShapes)
	canvas.GET("/shapes/:shape_id", handler.GetShape)
	canvas.PUT("/shapes/:shape_id", handler.UpdateShape)
	canvas.DELETE("/shapes/:shape_id", handler.DeleteShape)
}
