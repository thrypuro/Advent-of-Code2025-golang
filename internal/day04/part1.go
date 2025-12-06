package day04

import (
	"bufio"
	"log"
	"os"
)

func calcAdajacentSum(grid [][]int) int {
	adj_sum := 0
	length := len(grid)
	for i := range grid {
		width := len(grid[i])
		for j := range grid[i] {
			elem := grid[i][j]
			if elem != 1 {
				continue
			}
			total := 0
			adj_ind := []int{-1, 0, 1}
			for _, v := range adj_ind {
				for _, w := range adj_ind {

					v_t := v + i
					w_t := w + j
					if v_t == i && w_t == j {
						continue
					}
					if (v_t >= 0 && v_t < length) && (w_t >= 0 && w_t < width) {
						total += grid[v_t][w_t]
					}
				}
			}
			if total < 4 {
				adj_sum += 1
			}
		}
	}
	return adj_sum
}

func Part1() int {
	file, err := os.Open("internal/day04/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scan := bufio.NewScanner(file)
	grid := make([][]int, 0)
	for scan.Scan() {
		line := scan.Text()
		row := make([]int, 0)
		for v := range line {
			switch string(line[v]) {
			case ".":
				row = append(row, 0)
			case "@":
				row = append(row, 1)

			}

		}
		grid = append(grid, row)
	}
	// For debugging.
	// for i := range grid {
	// 	for j := range grid[i] {
	// 		print(grid[i][j])
	// 	}
	// 	println()
	// }
	sum := calcAdajacentSum(grid)
	return sum
}
