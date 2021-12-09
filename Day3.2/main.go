package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./Day3.2/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	lines := make([]string, 0, 1000)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	oxygenGeneratorRatingBinary, oxygenGeneratorRating, err := rating(lines, "gt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	co2ScrubberRatingBinary, co2ScrubberRating, err := rating(lines, "lt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf(
		"The result of Oxygen Generator Rating (%s, %d) * CO2 Scrubber Rating (%s, %d) is: %d\n",
		oxygenGeneratorRatingBinary,
		oxygenGeneratorRating,
		co2ScrubberRatingBinary,
		co2ScrubberRating,
		oxygenGeneratorRating*co2ScrubberRating,
	)
}

func rating(feed []string, condition string) (string, int, error) {
	lines := feed[0:]
	var pos int

	for len(lines) > 1 {
		stats := processLines(lines, pos)

		res := make([]string, 0)

		for _, l := range lines {
			switch {
			case stats[0] > stats[1] && condition == "gt" && l[pos] == '0':
				res = append(res, l)
			case stats[0] < stats[1] && condition == "lt" && l[pos] == '0':
				res = append(res, l)
			case stats[1] == stats[0] && condition == "lt" && l[pos] == '0':
				res = append(res, l)
			case stats[1] > stats[0] && condition == "gt" && l[pos] == '1':
				res = append(res, l)
			case stats[1] < stats[0] && condition == "lt" && l[pos] == '1':
				res = append(res, l)
			case stats[1] == stats[0] && condition == "gt" && l[pos] == '1':
				res = append(res, l)
			}
		}

		lines = res[0:]

		pos++
	}

	rating, err := strconv.ParseInt(lines[0], 2, 64)
	if err != nil {
		return "", 0, err
	}

	return lines[0], int(rating), nil
}

func processLines(lines []string, pos int) []int {
	stats := make([]int, 2)

	for _, l := range lines {
		switch l[pos] {
		case '0':
			stats[0]++
		case '1':
			stats[1]++
		}
	}

	return stats
}
