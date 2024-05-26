package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"

	"github.com/llgcode/draw2d/draw2dimg"
)

func printCallSheet(tiles []tile) {
	callPage := 0
	x0 := 0
	y0 := 0
	var sheet *image.RGBA = nil
	var gc *draw2dimg.GraphicContext = nil
	for i := range len(tiles) {
		t := tiles[i]
		if t.img == nil {
			continue
		}
		x1 := x0 + iSize
		if x1 > width {
			x0 = 0
			x1 = x0 + iSize
			y0 += iSize
		}
		y1 := y0 + iSize
		if y1 > height {
			if err := draw2dimg.SaveToPngFile(fmt.Sprintf("call%d.png", callPage), sheet); err != nil {
				log.Printf("Saving %q failed: %v", sheet, err)
			}
			sheet = nil
			y0 = 0
			y1 = y0 + iSize
			callPage += 1
		}
		if sheet == nil {
			sheet = image.NewRGBA(image.Rectangle{Max: image.Point{X: width, Y: height}})
			gc = draw2dimg.NewGraphicContext(sheet)
			gc.SetStrokeColor(color.Black)
			gc.SetFillColor(color.Transparent)
			gc.SetLineWidth(5)
		}

		gc.MoveTo(float64(x0), float64(y0))
		gc.LineTo(float64(x0), float64(y1))
		gc.LineTo(float64(x1), float64(y1))
		gc.LineTo(float64(x1), float64(y0))
		gc.Close()
		gc.Stroke()
		draw.Draw(sheet, image.Rect(int(x0), int(y0), int(x1), int(y1)), t.img, image.Point{}, draw.Over)
		x0 = x1
	}
	if err := draw2dimg.SaveToPngFile(fmt.Sprintf("call%d.png", callPage), sheet); err != nil {
		log.Printf("Saving %q failed: %v", sheet, err)
	}
}
