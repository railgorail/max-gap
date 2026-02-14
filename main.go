package main

import "math"

func maximumGap(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	mn, mx := nums[0], nums[0]
	for _, v := range nums {
		mn = min(mn, v)
		mx = max(mx, v)
	}

	// если все элементы равны, максимальная разница = 0
	if mn == mx {
		return 0
	}

	sz := max(1, (mx-mn)/(n-1))
	cnt := (mx-mn)/sz + 1

	b := make([][]int, cnt)
	for i := range b {
		b[i] = []int{math.MaxInt64, math.MinInt64}
	}

	for _, v := range nums {
		idx := (v - mn) / sz
		b[idx][0] = min(b[idx][0], v)
		b[idx][1] = max(b[idx][1], v)
	}

	prev := mn
	ans := 0
	for _, bucket := range b {
		if bucket[0] == math.MaxInt64 {
			continue // Пропускаем пустые бакеты
		}
		ans = max(ans, bucket[0]-prev)
		prev = bucket[1]
	}

	return ans
}

func main() {
	// Пример использования
	nums := []int{3, 6, 9, 1}
	result := maximumGap(nums)
	println(result) // Output: 3
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
