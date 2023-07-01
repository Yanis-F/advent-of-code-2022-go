package main

import "fmt"


func main() {
	slice := make([]int, 0)

	printSlice(slice)

	for i := 0; i < 1000; i++ {
		slice = append(slice, i)
		printSlice(slice)
	}
}

func printSlice(slice []int) {
	fmt.Printf("len: %d, cap: %d, content: %v\n", len(slice), cap(slice), slice)
}
