package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)


func getSums(input io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(input)

	var result []int
	var current_sum int

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			number, err := strconv.Atoi(line)
			if err != nil {
				return nil, err
			}
			current_sum += number
		} else {
			result = append(result, current_sum)
			current_sum = 0
		}
	}

	return result, nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	sums, err := getSums(file)
	if err != nil {
		log.Fatalln(err)
	}

	sort.Ints(sums)
	top3 := sums[len(sums)-3:]

	sumTop3 := top3[0] + top3[1] + top3[2]

	fmt.Println(sumTop3)
}
