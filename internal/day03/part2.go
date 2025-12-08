package day03

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

var seen_set = make(map[int]bool)

func findLargest(line string, last_char int, num string, maxLength int) int {
	if maxLength == 0 {
		num_int, _ := strconv.Atoi(num)
		if _, exists := seen_set[num_int]; !exists {
			seen_set[num_int] = true
			return num_int
		}
	}
	num_list := make([]int, 0)
	index_list := make(map[int][]int)

	for i := last_char; i < len(line); i++ {
		next_char := line[i]
		// println("Next Char ", string(next_char))
		potential := len(line)-i >= maxLength
		if potential {
			new_num := num + string(next_char)
			num_int, _ := strconv.Atoi(new_num)
			num_list = append(num_list, num_int)
			index_list[num_int] = append(index_list[num_int], i)
		}
	}
	// println("Last Char ", num)
	// // print the num_list
	// for i := range num_list {
	// 	println("Num List ", num_list[i])
	// }
	// println("Num List Length ", len(num_list))

	// get the max from num_list
	lol := max(num_list)
	smallest := 0
	for i := range index_list[lol] {
		candidate := index_list[lol][i]
		if candidate <= smallest {
			continue
		}
		smallest = candidate
	}

	// println("Max is ", max(num_list))
	return findLargest(line, smallest+1, num+strconv.Itoa(lol), maxLength-1)
}

// max returns the largest integer in a slice.
// If the slice is empty, it returns 0.
func max(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxVal := nums[0]
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

func Part2() int {
	file, err := os.Open("internal/day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scan := bufio.NewScanner(file)
	sum := 0
	for scan.Scan() {
		line := scan.Text()
		num := findLargest(line, 0, "", 12)
		println("Largest in Line is ", line, num)
		sum += num
	}
	return sum
}
