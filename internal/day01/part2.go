package day01

import (
	"bufio"
	"log"
	"os"
)

func mod(a, b int64) int64 {
	return (a%b + b) % b
}

func Part2() int {
	file, err := os.Open("internal/day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scan := bufio.NewScanner(file)
	var start int64 = 50
	var zeroes int64 = 0

	for scan.Scan() {
		line := scan.Text()
		var sym int64 = 1
		var a = len(line)
		if line[0] == 'L' {
			sym = -1
		}
		var numberLength int64 = 0
		for i := range a - 1 {
			numberLength = numberLength*10 + int64(line[i+1]-'0')
		}
		var offset = sym * numberLength
		var newStart = start + offset
		if newStart == 0 {
			zeroes += 1

		} else if newStart >= 100 {
			var localZero = newStart / 100
			zeroes += localZero

		} else if newStart < 0 {
			var localZero = (-newStart + 100) / 100
			zeroes += localZero
			if start == 0 {
				zeroes -= 1
			}
		}
		start = mod(newStart, 100)
	}
	return int(zeroes)
}
