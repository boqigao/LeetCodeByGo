package lc135

func candy(ratings []int) int {

	length := len(ratings)
	leftToRight := make([]int, length)
	rightToLeft := make([]int, length)

	for i := 0; i < length; i++ {
		leftToRight[i] = 1
		rightToLeft[i] = 1
	}

	for i := 1; i < length; i++ {
		if ratings[i] > ratings[i-1] {
			leftToRight[i] = leftToRight[i-1] + 1
		}
	}

	for i := length - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			rightToLeft[i] = rightToLeft[i+1] + 1
		}
	}

	sum := 0

	for i := 0; i < length; i++ {
		sum += max(rightToLeft[i], leftToRight[i])
	}
	return sum

}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
