package models

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
	"strconv"
	"strings"
)

const IMAGE_NAME = "canvas.png"

func CreateImage(shapes []Shape) (string, error) {
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
			drawTriangle(canvas, x, y, width, height, sideLeft, sideRight, sideBase, convertHexToRGBA(color))

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
	// มุมฉาก
	ff := 0
	ff2 := 0
	count := 0
	for iii := x; iii <= sideBase; iii++ {
		// if (sideBase - sideLeft) < iii {
		ff2++
		img.Set(sideBase-iii+x, iii, c)
		// }
		// img.Set(sideBase-iii-x, iii+x, c)

		// if (sideBase - sideRight) < iii { // ขวา
		ff++
		img.Set(sideBase, iii, c)

		// }

		img.Set(iii, sideBase, c)
		count = iii
	}

	log.Println("count", count)
	log.Println("ff", ff)
	log.Println("ff2", ff2)

	// for i := y; i < count; i++ {
	// 	if (sideBase - sideLeft) < i {
	// 		img.Set(sideBase-i+y, sideRight+i, c)
	// 	}
	// }

	// draw triangle left to right
	// r := count
	// for i := y; i < count; i++ {
	// 	r--
	// 	if (sideBase - sideLeft) < i {

	// 		img.Set(sideBase-i+y, sideRight+i+r, c)
	// 	}
	// }

	// for i := 0; i < height; i++ {
	// 	for j := 0; j < width; j++ {
	// 		if i < sideLeft {
	// 			img.Set(x+j, y+i, c)
	// 		}
	// 		if i < sideRight {
	// 			img.Set(x+width-j, y+i, c)
	// 		}
	// 		if i < sideBase {
	// 			img.Set(x+width/2+j, y+i, c)
	// 		}
	// 	}
	// }
}
