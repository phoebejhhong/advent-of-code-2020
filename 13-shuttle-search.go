package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Bus struct {
	number int
	timing int
}

func getWaitTime(startTime int, bus int) int {
	remainder := startTime % bus
	return bus - remainder
}

func getPart1(startTime int, buses map[int]int) int {
	minWaitTime := -1
	minWaitBus := 0
	for bus, _ := range buses {
		waitTime := getWaitTime(startTime, bus)
		if minWaitTime == -1 || waitTime < minWaitTime {
			minWaitTime = waitTime
			minWaitBus = bus
		}
	}
	return minWaitBus * minWaitTime
}

func getPart2(buses map[int]int) int {
	startTime := 0
	step := 1
	for bus, timing := range buses {
		for (startTime+timing)%bus != 0 {
			startTime += step
		}
		step *= bus
	}

	return startTime
}

func main() {
	data, _ := ioutil.ReadFile("./inputs/13-input.txt")
	lines := strings.Split(string(data), "\n")
	startTime, _ := strconv.Atoi(lines[0])
	buses := map[int]int{}
	for i, busNumber := range strings.Split(lines[1], ",") {
		if busNumber != "x" {
			intNum, _ := strconv.Atoi(busNumber)
			buses[intNum] = i
		}
	}

	part1 := getPart1(startTime, buses)
	part2 := getPart2(buses)

	fmt.Println("📝 Part 1: ", part1)
	fmt.Println("📝 Part 2: ", part2)
}
