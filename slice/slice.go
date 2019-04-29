package main

// Sum all of numbers in the slice
func Sum(numbers []int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

// SumAll will take a varying number of slices, returning a new slice containing the totals for each slice passed in
func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for _, numbers := range numbersToSum {
		// sums[_] = Sum(numbers)
    sums = append(sums, Sum(numbers))
	}
	return sums
}

// SumAllTails returns the sum of all tails(everything but the firtst element) of each slice
func SumAllTails(numbersToSum ...[]int) []int {
    var sums []int
    for _, numbers := range numbersToSum {
        if len(numbers) == 0 {
            sums = append(sums, 0)
        } else {
            tail := numbers[1:]
            sums = append(sums, Sum(tail))
        }
    }
    return sums
}
