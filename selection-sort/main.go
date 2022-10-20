package main

import "fmt"

func main() {
	fmt.Println(selectionSort([]int{5, 3, 2, 6, 10}))
}

func selectionSort(arr []int) (newArr []int) {
	newArr = make([]int, 0, len(arr))

	for range arr {
		smallestIndex := findSmallest(arr)
		newArr = append(newArr, arr[smallestIndex])
		arr = removeElementAt(arr, smallestIndex)
	}

	return
}

func findSmallest(arr []int) (smallestIndex int) {
	if len(arr) == 0 {
		smallestIndex = -1
		return
	}

	for i, item := range arr {
		if item < arr[smallestIndex] {
			smallestIndex = i
		}
	}

	return
}

func removeElementAt(arr []int, index int) (newArr []int) {
	if len(arr) == 0 {
		newArr = make([]int, 0)
		return
	}

	if index >= len(arr) {
		newArr = arr
		return
	}

	newArr = make([]int, 0, len(arr)-1)

	for i, item := range arr {
		if i == index {
			continue
		}

		newArr = append(newArr, item)
	}

	return
}
