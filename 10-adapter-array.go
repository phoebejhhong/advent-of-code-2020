package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("./inputs/10-input.txt")
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]
	adapters := []int{}

	for _, stringAdapter := range lines {
		intAdapter, _ := strconv.Atoi(stringAdapter)
		adapters = append(adapters, intAdapter)
	}

	sort.Ints(adapters)
	prev := 0
	differenceMap := map[int]int{}
	adapterIndexMap := map[int]int{}

	for i, joltage := range adapters {
		diff := joltage - prev
		differenceMap[diff] += 1
		prev = joltage
		adapterIndexMap[joltage] = i
	}

	differenceMap[3] += 1
	fmt.Printf(" 📝 Number of 1-jolt difference is %v and number of 3-joit difference is %v. Multiplied? %v \n", differenceMap[1], differenceMap[3], differenceMap[1]*differenceMap[3])

	savedCounts := make([]int, len(adapters))
	savedCounts[len(adapters)-1] = 1
	for i := len(adapters) - 2; i >= 0; i-- {
		currentCount := 0
		currentVoltage := adapters[i]
		for diff := 1; diff <= 3; diff++ {
			if diffIndex, ok := adapterIndexMap[currentVoltage+diff]; ok {
				currentCount += savedCounts[diffIndex]
			}
		}
		savedCounts[i] = currentCount
	}
	result := 0
	for v := 1; v <= 3; v++ {
		if i, ok := adapterIndexMap[v]; ok {
			result += savedCounts[i]
		}
	}
	fmt.Printf(" 📝 The number of distinct ways to arrange the adapters is %v \n", result)

}
