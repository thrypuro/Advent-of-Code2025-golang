package day06

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part2() int {
	file, err := os.Open("internal/day06/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	matrix := make([][]int, 0)
	operands := make([]string, 0)
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()
		lines := strings.Split(line, " ")
		is_operand := false
		row := make([]int, 0)
		for i := range lines {
			if lines[i] == "" {
				continue
			} else if lines[i] == "*" || lines[i] == "+" {
				operands = append(operands, lines[i])
				is_operand = true
			} else {
				num, _ := strconv.Atoi(lines[i])
				row = append(row, num)
			}
		}
		if !is_operand {
			matrix = append(matrix, row)
		}
	}
	// Transpose the matrix
	if len(matrix) == 0 {
		return 0
	}
	transposed := make([][]int, len(matrix[0]))
	for i := range transposed {
		transposed[i] = make([]int, len(matrix))
		for j := range matrix {
			transposed[i][j] = matrix[j][i]
		}
	}
	transformed := make([][]int, len(matrix[0]))

	for i := range transformed {
		transformed[i] = make([]int, len(matrix))
		operand := operands[i]
		for j := range transposed[i] {

		}

	}
	return 0
}
