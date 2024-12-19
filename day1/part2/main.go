package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	var similarityScore int
	var line1 []int
	line2Map := make(map[int]int)

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
		line1 = append(line1, id1)

		id2, id2Err := strconv.Atoi(Ids[1])
		if id2Err != nil {
			panic(id2Err)
		}
		count, exists := line2Map[id2]
		if exists {
			line2Map[id2] = count + 1
		} else {
			line2Map[id2] = 1
		}
	}

	for i := 0; i < len(line1); i++ {
		count, exists := line2Map[line1[i]]
		if exists {
			similarityScore += line1[i] * count
		} else {
			similarityScore += 0
		}
	}
	println(similarityScore)
}
