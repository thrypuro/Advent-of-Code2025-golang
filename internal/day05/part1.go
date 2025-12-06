package day05

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	lower  int
	uppper int
}

func Part1() int {
	file, err := os.Open("internal/day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scan := bufio.NewScanner(file)
	id_range := make([]Range, 0)
	for scan.Scan() {
		line := scan.Text()
		if line == "" {
			break
		}
		split := strings.Split(line, "-")
		split_l, _ := strconv.Atoi(split[0])
		split_r, _ := strconv.Atoi(split[1])
		id_range = append(id_range, Range{split_l, split_r})
	}

	ids := make([]int, 0)
	for scan.Scan() {
		line := scan.Text()
		num, _ := strconv.Atoi(line)
		ids = append(ids, num)
	}

	// println("Ranges are ")
	// for i := range id_range {
	// 	println(id_range[i].lower, " - ", id_range[i].uppper)
	// }
	// println("Nums are")
	// for i := range ids {
	// 	println(ids[i])
	// }
	seen := make(map[int]bool)
	fresh := 0
	for i := range ids {
		id := ids[i]
		for j := range id_range {
			if id >= id_range[j].lower && id <= id_range[j].uppper && !seen[id] {
				// println("Id ", id, " is fresh")
				fresh += 1
				seen[id] = true
			}
		}
	}
	return fresh
}
