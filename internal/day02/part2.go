package day02

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func isValid2(num int) bool {
	num_seq := ""
	temp := strconv.Itoa(num)
	le := len(temp)
	for i := range le {
		num_seq = num_seq + string(temp[i])
		j := i + 1
		isPat := true
		slice := temp[j:]
		if slice == "" || len(num_seq) > len(slice) {
			return false
		}
		var k, gg int
		for k < le-j {
			gg = (k % len(num_seq))
			if slice[k] != num_seq[gg] {
				isPat = false
				break
			}
			k++
		}
		if isPat && gg == len(num_seq)-1 {
			return true
		}

	}
	return false
}

func Part2() int {
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
			if isValid2(j) {
				sum += j
			}
		}

	}
	return sum
}
