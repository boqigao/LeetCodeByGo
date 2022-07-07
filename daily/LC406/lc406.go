package main

import "sort"

func reconstructQueue(people [][]int) [][]int {
	// go里面的原则是，如果是true的话，a就排在b的前面
	sort.Slice(people, func(a, b int) bool {
		if people[a][0] == people[b][0] {
			// 如果 people[a][1] > people[b][1]， 那么a排在b的前面
			return people[a][1] > people[b][1]
		}
		// 如果people[a][0] < people[b][0]，那么a排在b前面
		return people[a][0] < people[b][0]
	})

	index := make([]int, len(people))
	ans := make([][]int, len(people))

	for i := range index {
		index[i] = i
	}

	for _, v := range people {
		ans[index[v[1]]] = v
		// insert的标准写法
		// s如果使用s...符号解压缩切片，则可以将切片直接传递给可变参数函数。在这种情况下，不会创建新的切片。
		// 相当于把后者的所有都打散，装进切片
		index = append(index[:v[1]], index[v[1]+1:]...)
	}
	return ans
}

func reconstructQueue1(people [][]int) [][]int {
	// go里面的原则是，如果是true的话，a就排在b的前面
	sort.Slice(people, func(a, b int) bool {
		if people[a][0] == people[b][0] {
			//在[a][0] == [b][0]的情况下
			//如果 people[a][1] < people[b][1]， 那么a排在b的前面
			return people[a][1] < people[b][1]
		}
		// 如果people[a][0] > people[b][0]，那么a排在b前面
		return people[a][0] > people[b][0]
	})

	ans := make([][]int, len(people))

	for _, v := range people {
		// insert的标准写法
		// s如果使用s...符号解压缩切片，则可以将切片直接传递给可变参数函数。在这种情况下，不会创建新的切片。
		// 相当于把后者的所有都打散，装进切片
		ans = append(ans[:v[1]+1], ans[v[1]:]...)
		ans[v[1]] = v
	}
	return ans[0:len(people)]
}

func main() {
	people := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	reconstructQueue1(people[:])
}
