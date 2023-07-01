package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)


func getMaxSum(input io.Reader) (int, error) {
	scanner := bufio.NewScanner(input)

	var max_sum, current_sum int

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			number, err := strconv.Atoi(line)
			if err != nil {
				return 0, err
			}
			current_sum += number
		} else {
			if current_sum > max_sum {
				max_sum = current_sum
			}
			current_sum = 0
		}
	}

	return max_sum, nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	result, err := getMaxSum(file)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(result)
}
