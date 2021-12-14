package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Had to get help on this one.  This problem gets me every year.
// https://github.com/lynerist/Advent-of-code-2021-golang/blob/master/Day_06/day06_b.go

func main() {
	file, err := os.Open("./Day6.2/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	fishes := scanner.Text()
	var lanternFishByAge [9]int

	for _, fish := range strings.Split(fishes, ",") {
		days, _ := strconv.Atoi(fish)
		lanternFishByAge[days]++
	}

	for evolution := 0; evolution < 256; evolution++ {
		justBred := lanternFishByAge[0]

		for days := range lanternFishByAge[:len(lanternFishByAge)-1] {
			lanternFishByAge[days] = lanternFishByAge[days+1]
		}

		lanternFishByAge[6] += justBred
		lanternFishByAge[8] = justBred
	}

	var fishCount int
	for _, fishes := range lanternFishByAge {
		fishCount += fishes
	}

	fmt.Println(fishCount)
}
