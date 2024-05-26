package main

import (
	"errors"
	"fmt"
	"image"
	_ "image/jpeg"
	"io/fs"
	"os"
	"path/filepath"

	"golang.org/x/image/draw"
)

func discover(paths []string) []tile {
	tiles := make([]tile, 0)
	if len(paths) > 0 {
		tiles = find(paths[0], tiles)
	} else {
		tiles = find("tiles", tiles)
	}
	return tiles
}

func find(path string, tiles []tile) []tile {
	s, err := os.Stat(path)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	if s.IsDir() {
		fs.WalkDir(os.DirFS("."), path, func(path string, d fs.DirEntry, err error) error {
			if err == nil {
				if tile, err := newTile(path, d); err == nil {
					tiles = append(tiles, tile)
				}
			}
			return nil
		})
	} else {
		tiles = append(tiles, tile{"", path, nil})
	}
	return tiles
}

type tile struct {
	path  string
	label string
	img   image.Image
}

func newTile(path string, d fs.DirEntry) (tile, error) {
	if d.Type().IsRegular() {
		ext := filepath.Ext(path)
		name := d.Name()
		label := name[:len(name)-len(ext)]
		f, _ := os.Open(path)
		defer f.Close()
		img, _, err := image.Decode(f)
		if err == nil {
			w := float64(img.Bounds().Dx())
			h := float64(img.Bounds().Dy())
			ratio := w / h
			tileImg := image.NewRGBA(image.Rect(0, 0, iSize, iSize))
			if ratio < 1 {
				w := (fSize - fSize*ratio) / 2
				draw.ApproxBiLinear.Scale(tileImg, image.Rect(int(w), 0, int(fSize-w), iSize), img, img.Bounds(), draw.Over, nil)
			} else {
				h := (fSize - fSize/ratio) / 2
				draw.ApproxBiLinear.Scale(tileImg, image.Rect(0, int(h), iSize, int(fSize-h)), img, img.Bounds(), draw.Over, nil)
			}
			return tile{path, label, tileImg}, nil
		} else {
			fmt.Println(d.Name(), err)
			return tile{path, label, img}, nil
		}
	}
	return tile{}, errors.New("unhandled type")
}
