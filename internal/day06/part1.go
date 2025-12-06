package day06

import (
	"bufio"
	"log"
	"os"
)

func Part1() int {
	file, err := os.Open("internal/day06/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()
		println(line)
	}
	return 0
}
