package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var preambleCount = 25

func checkIsSum(target int, fromNumbers []int) bool {
	if len(fromNumbers) < 2 {
		return true
	}
	for i := 0; i < len(fromNumbers)-1; i++ {
		for j := 1; j < len(fromNumbers); j++ {
			if i != j && (fromNumbers[i]+fromNumbers[j]) == target {
				return true
			}
		}
	}
	return false
}

func getViolatingNumbers(numberStrings []string) []int {
	violatingNumbers := []int{}
	numbersInConsideration := []int{}
	for i, numberString := range numberStrings {
		number, _ := strconv.Atoi(numberString)
		numbersInConsideration = append(numbersInConsideration, number)
		if i > preambleCount-1 {
			if !checkIsSum(number, numbersInConsideration) {
				violatingNumbers = append(violatingNumbers, number)
			}

			numbersInConsideration = numbersInConsideration[1:]
		}
	}

	return violatingNumbers
}

func findContiguousSetofSum(target int, fromNumbers []string) []int {
	for i := 0; i < len(fromNumbers)-1; i++ {
		for j := i + 1; j < len(fromNumbers); j++ {
			if i != j {
				thisSet := []int{}
				thisSetSum := 0
				for k := i; k < j; k++ {
					number, _ := strconv.Atoi(fromNumbers[k])
					thisSet = append(thisSet, number)
					thisSetSum += number
				}
				if thisSetSum == target {
					return thisSet
				}
			}
		}
	}
	return []int{}
}

func main() {
	data, _ := ioutil.ReadFile("./inputs/09-input.txt")
	lines := strings.Split(string(data), "\n")
	numbers := lines[:len(lines)-1]
	violatingNumbers := getViolatingNumbers(numbers)
	set := findContiguousSetofSum(violatingNumbers[0], numbers)
	var max, min int
	for _, numInSet := range set {
		if numInSet > max {
			max = numInSet
		}
		if min == 0 || numInSet < min {
			min = numInSet
		}
	}
	fmt.Println(" 📝 First number that does not follow the rule is", violatingNumbers[0])
	fmt.Println(" 📝 Encryption weakness is", max+min)
}
