package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func countAnyQuestions(input string) int {
	questionMap := map[string]bool{}
	uniqueQuestionCount := 0
	for _, rune := range input {
		char := string(rune)
		if !questionMap[char] {
			uniqueQuestionCount++
			questionMap[char] = true
		}
	}
	return uniqueQuestionCount
}

func countAllQuestions(input string, peopleCount int) int {
	questionMap := map[string]int{}
	for _, rune := range input {
		char := string(rune)
		questionMap[char] += 1
	}

	allAnsweredCount := 0
	for _, count := range questionMap {
		if count == peopleCount {
			allAnsweredCount++
		}
	}
	return allAnsweredCount
}

func main() {
	data, _ := ioutil.ReadFile("./06-input.txt")
	lines := strings.Split(string(data), "\n")
	currentGroupInput := ""
	currentGroupPeopleCount := 0
	anySum := 0
	allSum := 0

	for _, line := range lines {
		if line != "" {
			currentGroupPeopleCount++
			currentGroupInput += line
		} else {
			anySum += countAnyQuestions(currentGroupInput)
			allSum += countAllQuestions(currentGroupInput, currentGroupPeopleCount)
			currentGroupInput = ""
			currentGroupPeopleCount = 0
		}
	}

	fmt.Println("📝 Sum of any question answered yes in groups is ", anySum)
	fmt.Println("📝 Sum of all question answered yes in groups is ", allSum)
}
