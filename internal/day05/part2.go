package day05

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func skimIds(id_range []Range) []Range {

	sort.Slice(id_range, func(i, j int) bool {
		return id_range[i].lower < id_range[j].lower
	})

	merged := make([]Range, 0)

	for _, r := range id_range {
		if len(merged) == 0 || r.lower > merged[len(merged)-1].uppper {
			merged = append(merged, Range{r.lower, r.uppper})
		} else {
			if r.uppper > merged[len(merged)-1].uppper {
				merged[len(merged)-1].uppper = r.uppper
			}
		}
	}
	return merged
}

func Part2() int {
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
	reduce_range := skimIds(id_range)
	sum := 0
	for i := range reduce_range {
		sum += (reduce_range[i].uppper - reduce_range[i].lower) + 1
	}
	return sum
}
