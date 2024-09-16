package iteration

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numListsToSum ...[]int) (sums []int) {
	sums = []int{}
	for _, numbers := range numListsToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}

func SumAllTail(numListsToSum ...[]int) []int {
	sum := make([]int, len(numListsToSum))
	for i, list := range numListsToSum {
		if len(list) == 0 {
			sum[i] = 0
			continue
		}
		sum[i] = Sum(list[1:])
	}
	return sum
}
