package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
)

var pages int
var output = flag.String("o", "bingo.pdf", "Output file")
var bingo_tiles []tile

func main() {
	flag.IntVar(&pages, "p", 20, "Pages to generate")
	flag.IntVar(&pages, "pages", 20, "Pages to generate")
	flag.Parse()
	tiles := discover(flag.Args())
	fmt.Println(tiles)
}

func discover(paths []string) []tile {
	tiles := make([]tile, 0)
	if len(paths) > 0 {
		find(paths[0], &tiles)
	} else {
		find("tiles", &tiles)
	}
	return tiles
}

func find(path string, tiles *[]tile) {
	s, err := os.Stat(path)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	if s.IsDir() {
		fs.WalkDir(os.DirFS(path), path, func(path string, d fs.DirEntry, err error) error {
			if tile, err := newTile(d); err == nil {
				*tiles = append(*tiles, tile)
			}
			return nil
		})
	} else {
		*tiles = append(*tiles, tile{"", path})
	}
}

type tile struct {
	image string
	label string
}

func newTile(d fs.DirEntry) (tile, error) {
	if d.Type().IsRegular() {
		return tile{"", d.Name()}, nil
	}
	return tile{}, errors.New("Unhandled type")
}
