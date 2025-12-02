package day02

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func isValid(num int) bool {
	num_seq := ""
	temp := num
	for temp > 0 {
		num_seq = strconv.Itoa(int(temp%10)) + num_seq
		temp = temp / 10
		if strconv.Itoa(int(temp)) == num_seq {
			return true
		}
	}
	return false
}

func Part1() int {
	file, err := os.Open("internal/day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scan := bufio.NewScanner(file)
	var line_sep []string
	for scan.Scan() {
		line := scan.Text()
		line_sep = strings.Split(line, ",")
	}
	var sum int
	for i := range len(line_sep) {
		split_rang := strings.Split(line_sep[i], "-")
		lower_str := (split_rang[0])
		upper_str := (split_rang[1])
		lower, _ := strconv.Atoi(lower_str)
		upper, _ := strconv.Atoi(upper_str)

		for j := lower; j <= upper; j++ {
			if isValid(j) {
				sum += j
			}
		}

	}
	return sum
}
