package main

import (
	"canvas/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	_ = r.Run(":7001")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	canvasRepo := controllers.NewCanvasTable()
	r.POST("/canvas", canvasRepo.CreateCanvas)
	r.GET("/canvas", canvasRepo.GetCanvas)
	r.GET("/canvas/:canvas_id", canvasRepo.GetCanvasById)
	r.PUT("/canvas/:canvas_id", canvasRepo.UpdateCanvas)
	r.DELETE("/canvas/:canvas_id", canvasRepo.DeleteCanvas)

	shapeRepo := controllers.NewShapeTable()
	r.POST("/canvas/:canvas_id/shape/:shape_type", shapeRepo.CreateShape)
	r.GET("/canvas/:canvas_id/shapes", shapeRepo.GetShape)
	r.GET("/canvas/shapes/:shape_id", shapeRepo.GetShapeById)
	r.PUT("/canvas/shapes/:shape_id", shapeRepo.UpdateShape)
	r.DELETE("/canvas/shapes/:shape_id", shapeRepo.DeleteShape)
	r.GET("/canvas/:canvas_id/image", shapeRepo.GetImage)

	return r
}
