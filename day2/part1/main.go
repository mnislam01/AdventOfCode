package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, openErr := os.Open("input.txt")
	levelMatrix := make(map[int][]int)
	if openErr != nil {
		panic(openErr)
	}
	defer file.Close()
	safeCount := 0

	scanner := bufio.NewScanner(file)
	k := 0
	for scanner.Scan() {
		line := scanner.Text()
		levels := strings.Fields(line)
		var levelList []int
		for i := 0; i < len(levels); i++ {
			levelInt, _ := strconv.Atoi(levels[i])
			levelList = append(levelList, levelInt)
		}
		levelMatrix[k] = levelList
		k++
	}
	for j := 0; j < len(levelMatrix); j++ {
		levelsData := levelMatrix[j]
		increasing := false
		decreasing := false
		var safe = true
		for f := 1; f < len(levelsData); f++ {
			// Check if the values are in increasing order or decreasing order
			if levelsData[f-1] < levelsData[f] {
				if decreasing {
					// If previous values were in decreasing order
					// then the current values are not in increasing order
					safe = false
					continue
				}

				diff := levelsData[f] - levelsData[f-1]
				if diff < 1 || diff > 3 {
					safe = false
					continue
				}

				increasing = true
			} else if levelsData[f-1] > levelsData[f] {
				if increasing {
					// If previous values were in increasing order
					// then the current values are not in decreasing order
					safe = false
					continue
				}

				diff := levelsData[f-1] - levelsData[f]

				if diff < 1 || diff > 3 {
					safe = false
					continue
				}

				decreasing = true
			} else if levelsData[f-1] == levelsData[f] {
				// If the values are equal then they are not in increasing or decreasing order
				safe = false
				break
			}
		}

		if safe {
			safeCount++
		}
	}
	println(safeCount)
}
