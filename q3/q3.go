package q3

func FindMinMaxAverage(numbers []int) (int, int, float64, error) {
		min := numbers[0]
	max := numbers[0]
	sum := 0

	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		sum += num
	}

	average := float64(sum) / float64(len(numbers))

	return min, max, average, nil
}


