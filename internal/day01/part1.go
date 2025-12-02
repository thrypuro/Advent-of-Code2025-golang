package day01

import (
	"bufio"
	"log"
	"os"
)

func Part1() int {
	file, err := os.Open("internal/day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scan := bufio.NewScanner(file)
	var start uint = 50
	var zeroes = 0
	for scan.Scan() {
		line := scan.Text()
		var sym uint = 1
		var numberLength = len(line)
		if line[0] == 'L' {
			sym = 99
		}
		var parsedNum uint = 0
		for i := range numberLength - 1 {
			parsedNum = parsedNum*10 + uint(line[i+1]-'0')
		}
		start = (start + sym*parsedNum) % 100
		if start == 0 {
			zeroes += 1
		}
	}
	return zeroes
}
