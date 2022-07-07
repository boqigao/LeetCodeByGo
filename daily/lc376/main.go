package main

func wiggleMaxLength(nums []int) int {
	var length = len(nums)
	if len(nums) < 2 {
		return len(nums)
	}
	up := make([]int, length)
	down := make([]int, length)

	up[0] = 1
	down[0] = 1

	for i := 1; i < length; i++ {
		if nums[i] == nums[i-1] {
			up[i] = up[i-1]
			down[i] = down[i-1]
		} else if nums[i] < nums[i-1] {
			up[i] = up[i-1]
			down[i] = up[i-1] + 1
		} else {
			up[i] = down[i-1] + 1
			down[i] = down[i-1]
		}
	}

	if up[length-1] > down[length-1] {
		return up[length-1]
	} else {
		return down[length-1]
	}
}

func main() {
	arr := []int{1, 7, 4, 9, 2, 5}
	println(wiggleMaxLength(arr))
}
