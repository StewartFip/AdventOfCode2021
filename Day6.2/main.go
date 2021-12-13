package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./Day6.2/input_test.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(split)

	fishes := make([]int, 0)
	for scanner.Scan() {
		timer, _ := strconv.Atoi(scanner.Text())
		fishes = append(fishes, timer)
	}

	newFish := make([]int, 0)

	for day := 1; day <= 256; day++ {
		fishes = append(fishes, newFish...)
		newFish = make([]int, 0)

		for i := range fishes {
			if fishes[i] == 0 {
				fishes[i] = 6
			} else {
				fishes[i]--
			}

			if fishes[i] == 0 {
				newFish = append(newFish, 9)
			}
		}
	}

	fmt.Printf("Total fish after 80 days is %d\n", len(fishes))
}

func split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Return nothing if at end of file and no data passed
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// Find the index of the input where a comma is present.
	if i := strings.Index(string(data), ","); i >= 0 {
		return i + 1, data[0:i], nil
	}

	// If at end of file with data return the data
	if atEOF {
		return len(data), data, nil
	}

	return
}
