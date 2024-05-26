package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"
	"os"

	"github.com/llgcode/draw2d/draw2dimg"
	xdraw "golang.org/x/image/draw"
)

var fSize float64 = 292
var iSize = 292
var top float64 = 600
var width = 3000
var height = 2100
var right = 1540

func render(page int, one, two [][]tile) {
	sheet := image.NewRGBA(image.Rectangle{Max: image.Point{X: width, Y: height}})
	headerFile, err := os.Open("header.png")
	if err != nil {
		return
	}
	defer headerFile.Close()
	header, _, err := image.Decode(headerFile)
	if err != nil {
		return
	}
	headerImg := image.NewRGBA(image.Rect(0, 0, width/2, height))
	xdraw.CatmullRom.Scale(headerImg, headerImg.Bounds(), header, header.Bounds(), draw.Over, nil)

	gc := draw2dimg.NewGraphicContext(sheet)
	draw.Draw(sheet, image.Rect(0, 0, width/2-40, sheet.Bounds().Dy()-40), headerImg, image.Point{}, draw.Over)
	draw.Draw(sheet, image.Rect(right, 0, width, sheet.Bounds().Dy()-40), headerImg, image.Point{}, draw.Over)
	gc.SetStrokeColor(color.Black)
	gc.SetFillColor(color.Transparent)
	gc.SetLineWidth(1)
	gc.MoveTo(float64(width)/2, 0)
	gc.LineTo(float64(width)/2, float64(sheet.Bounds().Dy()))
	gc.SetLineWidth(5)

	// gc.SetFontData(draw2d.FontData{Family: draw2d.FontFamilyMono, Style: draw2d.FontStyleBold | draw2d.FontStyleItalic})
	gc.SetFontSize(14)
	drawSheet(one, 0, sheet, gc)
	drawSheet(two, float64(right), sheet, gc)
	if err := draw2dimg.SaveToPngFile(fmt.Sprintf("out%d.png", page), sheet); err != nil {
		log.Printf("Saving %q failed: %v", sheet, err)
	}
}

// var bingo = []string{"B", "I", "N", "G", "O"}

func drawSheet(tiles [][]tile, left float64, sheet *image.RGBA, gc *draw2dimg.GraphicContext) {
	for xx := range 5 {
		x := float64(xx)
		x0 := left + fSize*x
		x1 := x0 + fSize
		// gc.FillStringAt(bingo[xx], left+size*x, 50)
		for yy := range 5 {
			y := float64(yy)
			y0 := top + fSize*y
			y1 := y0 + fSize
			gc.MoveTo(x0, y0)
			gc.LineTo(x0, y1)
			gc.LineTo(x1, y1)
			gc.LineTo(x1, y0)
			gc.Close()
			gc.Stroke()

			t := tiles[xx][yy]
			if t.img != nil {
				draw.Draw(sheet, image.Rect(int(x0), int(y0), int(x1), int(y1)), t.img, image.Point{}, draw.Over)
			}
		}
	}
}
