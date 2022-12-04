package main

import "fmt"

func main() {
	fmt.Println(Sum([]int{2, 4, 6}))
	fmt.Println(QuickSort([]int{3, 5, 2, 1, 4}))
}

// Sum is a function that sums all slices using functional programming (divide & conquer)
func Sum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	return nums[0] + Sum(nums[1:])
}

// QuickSort is a function that implements divide & conquer to sorting the numbers
func QuickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	pivot := nums[0]

	less := Less(nums[1:], pivot)
	greater := Greater(nums[1:], pivot)

	result := append(QuickSort(less), pivot)
	result = append(result, QuickSort(greater)...)
	return result
}

func Less(nums []int, target int) []int {
	results := make([]int, 0)

	for _, num := range nums {
		if num <= target {
			results = append(results, num)
		}
	}

	return results
}

func Greater(nums []int, target int) []int {
	results := make([]int, 0)

	for _, num := range nums {
		if num > target {
			results = append(results, num)
		}
	}

	return results
}
