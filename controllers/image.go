package controllers

import (
	"canvas/models"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
	"strconv"
	"strings"
)

const IMAGE_NAME = "canvas.png"

func CreateImage(shapes []models.Shape) (string, error) {
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

		if shape.Type == "circle" {
			drawCircle(canvas, x, y, radius, convertHexToRGBA(color))
		}

		if shape.Type == "rectangle" {
			item := image.Rect(x, y, width, height)
			draw.Draw(canvas, item, &image.Uniform{convertHexToRGBA(color)}, image.ZP, draw.Src)
		}

		if shape.Type == "triangle" {
			drawTriangle(canvas, x, y, sideLeft, sideRight, sideBase, convertHexToRGBA(color))

		}

	}

	file, err := os.Create(IMAGE_NAME)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	jpeg.Encode(file, canvas, &jpeg.Options{Quality: 100})

	return IMAGE_NAME, nil
}

func convertHexToRGBA(hex string) color.RGBA {
	hex = strings.TrimPrefix(hex, "#")
	r, _ := strconv.ParseInt(hex, 16, 32)
	return color.RGBA{uint8(r >> 16), uint8((r & 0x00ff00) >> 8), uint8(r & 0x0000ff), 255}
}

func drawCircle(img draw.Image, x, y, radius int, c color.Color) {
	xc, yc, dx, dy := radius-1, 0, 1, 1
	err := dx - (radius * 2)
	for xc > yc {

		img.Set(x-xc, y+yc, c)
		img.Set(x-yc, y+xc, c)
		img.Set(x-yc, y+xc, c)
		img.Set(x-xc, y+yc, c)
		img.Set(x-xc, y-yc, c)
		img.Set(x-yc, y-xc, c)
		img.Set(x-yc, y-xc, c)
		img.Set(x-xc, y-yc, c)

		if err <= 0 {
			yc++
			err += dy
			dy += 2
		}
		if err > 0 {
			xc--
			dx += 2
			err += dx - (radius * 2)
		}
	}
}

func drawTriangle(img draw.Image, x, y, sideLeft, sideRight, sideBase int, c color.Color) {
	for i := y; i <= y+sideLeft; i++ {
		img.Set(sideLeft-i+y, sideRight+i, c)
		img.Set(sideLeft+i-y, sideRight+i, c)
	}

	for iii := x; iii <= sideBase*2; iii++ {
		img.Set(iii, sideBase, c)
	}
}
