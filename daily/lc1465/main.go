package main

import (
	"sort"
)

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)

	lenH := len(horizontalCuts)
	lenV := len(verticalCuts)

	maxHeight := max(horizontalCuts[0], h-horizontalCuts[lenH-1])
	for i := 1; i < lenH; i++ {
		maxHeight = max(maxHeight, horizontalCuts[i]-horizontalCuts[i-1])
	}

	maxWidth := max(verticalCuts[0], w-verticalCuts[lenV-1])

	for i := 1; i < lenV; i++ {
		maxWidth = max(maxWidth, verticalCuts[i]-verticalCuts[i-1])
	}

	return (maxWidth * maxHeight) % 1000000007
}
