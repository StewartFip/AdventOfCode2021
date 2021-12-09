package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	depths := make([]int, 0)

	file, err := os.Open("./Day1.1/input.txt")
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

	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			increased++
		}
	}

	fmt.Printf("Increased Depth Count: %d\n", increased)
}
