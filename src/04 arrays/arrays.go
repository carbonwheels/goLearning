package array

// Version 1: Sum calculates the total from an array of numbers.
func Sum(numbers [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return sum
}

//  Version 2: Sum calculates the total from an array of numbers.
func Sum2(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// Version 3: Sum calculates the total from a slice of numbers.
func Sum3(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
