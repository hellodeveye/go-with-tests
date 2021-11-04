package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	i := len(numbersToSum)
	sums := make([]int, i)
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}
