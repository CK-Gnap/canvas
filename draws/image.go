package models

import (
	models_interfaces "canvas/models/interfaces"
	"errors"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
	"strconv"
	"strings"
)

// func CreateImage(name string, shapes []Shape) (string, error) {
func CreateImage(name string, shapes []models_interfaces.ShapeInterface) (string, error) {

	canvas := image.NewRGBA(image.Rect(0, 0, 1000, 1000)) // ส่งเข้ามา
	bgColor := color.RGBA{255, 255, 255, 255}             // ส่งเข้ามา
	draw.Draw(canvas, canvas.Bounds(), &image.Uniform{bgColor}, image.ZP, draw.Src)

	for _, shape := range shapes {
		if shape.GetType() == TypeEnum(CIRCLE) {
			circle := shape.(*Circle)
			drawCircle(canvas, int(circle.X), int(circle.Y), int(circle.Radius), convertHexToRGBA(circle.Color))
		} else if shape.GetType() == TypeEnum(RECTANGLE) {
			rectangle := shape.(*Rectangle)
			item := image.Rect(int(rectangle.X), int(rectangle.Y), int(rectangle.Width), int(rectangle.Height))
			draw.Draw(canvas, item, &image.Uniform{convertHexToRGBA(rectangle.Color)}, image.ZP, draw.Src)
		} else if shape.GetType() == TypeEnum(TRIANGLE) {
			triangle := shape.(*Triangle)
			drawTriangle(canvas, int(triangle.X), int(triangle.Y), int(triangle.Width), int(triangle.Height), int(triangle.SideLeft), int(triangle.SideRight), int(triangle.SideBase), convertHexToRGBA(triangle.Color))
		} else {
			return "", errors.New("Invalid shape type")
		}
	}

	canvasName := strings.ReplaceAll(strings.ToLower(name), " ", "_")
	canvasName = "canvas_" + canvasName + ".jpg" // รับจากข้างนอกมาให้ใช้ได้กับอันอื่นด้วย
	file, err := os.Create(canvasName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	jpeg.Encode(file, canvas, &jpeg.Options{Quality: 100})

	return canvasName, nil
}

func convertHexToRGBA(hex string) color.RGBA {
	hex = strings.TrimPrefix(hex, "#")
	r, _ := strconv.ParseInt(hex, 16, 32)
	return color.RGBA{uint8(r >> 16), uint8((r & 0x00ff00) >> 8), uint8(r & 0x0000ff), 255}
}

func drawCircle(img draw.Image, x0, y0, r int, c color.Color) {
	x, y, dx, dy := r-1, 0, 1, 1
	err := dx - (r * 2)
	for x > y {

		img.Set(x0+x, y0+y, c)
		img.Set(x0+y, y0+x, c)
		img.Set(x0-y, y0+x, c)
		img.Set(x0-x, y0+y, c)
		img.Set(x0-x, y0-y, c)
		img.Set(x0-y, y0-x, c)
		img.Set(x0+y, y0-x, c)
		img.Set(x0+x, y0-y, c)

		if err <= 0 {
			y++
			err += dy
			dy += 2
		}
		if err > 0 {
			x--
			dx += 2
			err += dx - (r * 2)
		}
	}

}

func drawTriangle(img draw.Image, x, y, width, height, sideLeft, sideRight, sideBase int, c color.Color) { // ไม่เอา left, right, base
	for i := x; i <= sideBase; i++ {
		img.Set(sideBase-i+x, i, c)
		img.Set(sideBase, i, c)
		img.Set(i, sideBase, c)
	}
}
