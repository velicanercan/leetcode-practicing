package main

func maxArea(height []int) (result int) {
	length := len(height)

	var left, right = 0, length - 1

	for left < right {
		area := min(height[left], height[right]) * (right - left)
		result = max(result, area)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return
}
