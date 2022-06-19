package models

import (
	"canvasOld/models"
	"errors"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
	"strconv"
	"strings"
)

func CreateImage(name string, shapes []Shape) (string, error) {
	canvas := image.NewRGBA(image.Rect(0, 0, 1000, 1000))
	bgColor := color.RGBA{255, 255, 255, 255}
	draw.Draw(canvas, canvas.Bounds(), &image.Uniform{bgColor}, image.ZP, draw.Src)

	for _, shape := range shapes {
		color := shape.Color
		x := int(shape.X)
		y := int(shape.Y)
		width := int(shape.Width)
		height := int(shape.Height)
		radius := int(shape.Radius)
		sideLeft := int(shape.SideLeft)
		sideRight := int(shape.SideRight)
		sideBase := int(shape.SideBase)

		if shape.Type == TypeEnum(models.CIRCLE) {
			drawCircle(canvas, x, y, radius, convertHexToRGBA(color))
		} else if shape.Type == TypeEnum(models.RECTANGLE) {
			item := image.Rect(x, y, width, height)
			draw.Draw(canvas, item, &image.Uniform{convertHexToRGBA(color)}, image.ZP, draw.Src)
		} else if shape.Type == TypeEnum(models.TRIANGLE) {
			drawTriangle(canvas, x, y, width, height, sideLeft, sideRight, sideBase, convertHexToRGBA(color))
		} else {
			return "", errors.New("Invalid shape type")
		}

	}
	canvasName := strings.ReplaceAll(strings.ToLower(name), " ", "_")
	canvasName = "canvas_" + canvasName + ".jpg"
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

func drawTriangle(img draw.Image, x, y, width, height, sideLeft, sideRight, sideBase int, c color.Color) {
	for i := x; i <= sideBase; i++ {
		img.Set(sideBase-i+x, i, c)
		img.Set(sideBase, i, c)
		img.Set(i, sideBase, c)
	}
}
