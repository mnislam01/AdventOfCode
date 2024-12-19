package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var sumOfDifference int
	var line1 []int
	var line2 []int

	file, openErr := os.Open("input.txt")
	if openErr != nil {
		panic(openErr)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		Ids := strings.Fields(scanner.Text())
		id1, id1Err := strconv.Atoi(Ids[0])
		if id1Err != nil {
			panic(id1Err)
		}
		id2, id2Err := strconv.Atoi(Ids[1])
		if id2Err != nil {
			panic(id2Err)
		}
		line1 = append(line1, id1)
		line2 = append(line2, id2)
	}
	sort.Ints(line1)
	sort.Ints(line2)

	for i := 0; i < len(line1); i++ {
		diff := line2[i] - line1[i]
		sumOfDifference += abs(diff)
	}
	println(sumOfDifference)
}
