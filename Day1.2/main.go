package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	depths := make([]int, 0)

	file, err := os.Open("./Day1.2/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		depths = append(depths, depth)
	}

	var increased int

	for i := 0; i < len(depths)-3; i++ {
		if sumDepths(depths[i+1:i+4]) > sumDepths(depths[i:i+3]) {
			increased++
		}
	}

	fmt.Printf("Increased Depth Count: %d\n", increased)
}

func sumDepths(depths []int) int {
	var depth int

	for _, d := range depths {
		depth += d
	}

	return depth
}
