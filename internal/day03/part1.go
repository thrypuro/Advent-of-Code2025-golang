package day03

import (
	"bufio"
	"log"
	"os"
)

func Part1() int {
	file, err := os.Open("internal/day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scan := bufio.NewScanner(file)
	sum := 0
	for scan.Scan() {
		line := scan.Text()
		var largest, largest2 int
		for i := range len(line) - 1 {
			d_i := int(line[i] - '0')
			for j := i + 1; j < len(line); j++ {
				l_c := largest*10 + largest2
				d_j := int(line[j] - '0')
				d := d_i*10 + d_j
				if d > l_c {
					largest = d_i
					largest2 = d_j
				}
			}
		}
		sum += (largest*10 + largest2)

	}
	return sum
}
