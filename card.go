package main

func generateCard(c int, columns []column, free *tile) [][]tile {
	return [][]tile{
		generateColumn(c, free, &columns[0]),
		generateColumn(c, free, &columns[1]),
		generateColumn(c, free, &columns[4]),
		generateColumn(c, free, &columns[2]),
		generateColumn(c, free, &columns[3]),
	}
}

func generateColumn(c int, free *tile, col *column) []tile {
	membership := col.memberCombinations[c%col.comboLength]
	order := col.ordering[c%col.orderLength]
	data := make([]tile, 0, col.pick)
	for i := range col.pick {
		o := order[i]
		m := membership[o]
		data = append(data, col.tiles[m])
	}
	if col.pick == 4 {
		data = append(data, *free)
		data[2], data[4] = data[4], data[2]
	}
	return data
}
