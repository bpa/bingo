package main

import (
	"math/rand"
)

var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199}

func buildColumns(tiles []tile, knownRand *rand.Rand) []column {
	knownRand.Shuffle(len(tiles), func(i, j int) { tiles[i], tiles[j] = tiles[j], tiles[i] })
	columns := make([]column, 0, 5)
	full := len(tiles) % 5
	partial := 5 - full
	elements := len(tiles) / 5

	columns, start := addColumns(columns, tiles, full, elements+1, 0)
	columns, _ = addColumns(columns, tiles, partial, elements, start)
	columns[4].pick = 4
	next := primes[len(primes)-1]
	for i := range len(columns) {
		c := &columns[i]
		c.memberCombinations = generateCombinations(len(c.tiles), c.pick)
		shuffle(c.memberCombinations, knownRand)
		p := len(c.memberCombinations)
		next = pickPrime(min(next, p))
		c.comboLength = next
		c.ordering = permutations(c.pick)
		shuffle(c.ordering, knownRand)
		c.orderLength = pickPrime(len(c.ordering))
	}
	return columns
}

func addColumns(columns []column, tiles []tile, n, elements, start int) ([]column, int) {
	if n > 0 {
		for range n {
			columns = append(columns, column{tiles[start : start+elements], [][]int{}, [][]int{}, 5, 0, 0})
			start += elements
		}
	}
	return columns, start
}

func generateCombinations(n, pick int) [][]int {
	combos := fact(n) / (fact(n-pick) * fact(pick))
	combinations := make([][]int, 0, combos)
	entry := make([]int, 0, pick)
	return addCombinations(combinations, entry, n, pick)
}

func addCombinations(combos [][]int, entry []int, n, m int) [][]int {
	if n < m {
		return combos
	}
	if m == 0 {
		c := make([]int, len(entry))
		copy(c, entry)
		return append(combos, c)
	}
	combos = addCombinations(combos, entry, n-1, m)
	entry = append(entry, n-1)
	combos = addCombinations(combos, entry, n-1, m-1)
	return combos
}

// var factorials = make([]int, 0)
func fact(n int) int {
	if n == 1 {
		return 1
	}
	return n * fact(n-1)
}

func permutations(n int) [][]int {
	perms := make([][]int, 0, fact(n))
	entry := make([]int, 0, n)
	for i := range n {
		entry = append(entry, i)
	}
	return genMutations(n, entry, perms)
}

func genMutations(k int, entry []int, perms [][]int) [][]int {
	if k == 1 {
		next := make([]int, len(entry))
		copy(next, entry)
		return append(perms, next)
	}
	perms = genMutations(k-1, entry, perms)
	for i := range k - 1 {
		if k%2 == 0 {
			entry[i], entry[k-1] = entry[k-1], entry[i]
		} else {
			entry[0], entry[k-1] = entry[k-1], entry[0]
		}
		perms = genMutations(k-1, entry, perms)
	}
	return perms
}

func pickPrime(n int) int {
	for i := len(primes) - 1; i > 0; i -= 1 {
		if primes[i] < n {
			return primes[i]
		}
	}
	return 0
}

func shuffle(data [][]int, knownRand *rand.Rand) {
	knownRand.Shuffle(len(data), func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})
}

type column struct {
	tiles              []tile
	memberCombinations [][]int
	ordering           [][]int
	pick               int
	comboLength        int
	orderLength        int
}
