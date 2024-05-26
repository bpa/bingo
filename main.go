package main

import (
	"flag"
	"fmt"
	"math/rand"
	"slices"
)

var pages int

// var _output = flag.String("o", "bingo.pdf", "Output file")

func init() {
	flag.IntVar(&pages, "p", 0, "Pages to generate")
	// flag.IntVar(&pages, "pages", 25, "Pages to generate")
}

func main() {
	flag.Parse()
	knownRand := rand.New(rand.NewSource(0))
	tiles := discover(flag.Args())
	freeIdx := slices.IndexFunc(tiles, func(t tile) bool { return t.label == "Free" })
	free := tiles[freeIdx]
	tiles[freeIdx] = tiles[0]
	tiles = tiles[1:]
	columns := buildColumns(tiles, knownRand)
	fmt.Println(pages)
	for i := range pages {
		fmt.Println("Page", i)
		card1 := generateCard(i*2, columns, &free)
		card2 := generateCard(i*2+1, columns, &free)
		render(i, card1, card2)
	}
	printCallSheet(tiles)
}
