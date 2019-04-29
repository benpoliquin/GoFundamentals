package main

// Sum all of numbers in the slice
func Sum(numbers []int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}
