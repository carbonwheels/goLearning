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

// Version 4: SumAll calculates the respective sums of every slice passed in.
// Determines the length of each array and returns each sum
func SumAll4(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum3(numbers)
	}

	return sums
}

// Version 5: SumAll calculates the respective sums of every slice passed in.
// Loops through each array and returns each sum
func SumAll5(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum3(numbers))
	}

	return sums
}
