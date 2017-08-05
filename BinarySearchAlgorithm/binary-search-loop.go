/*
	Binary search implementation in Go using loop
*/
package main

//BinarySearchLoop test
func binarySearchLoop(array []int, target int, lowIndex int, highIndex int) int {
	startIndex := lowIndex
	endIndex := highIndex
	for startIndex < endIndex {
		mid := int((startIndex + endIndex) / 2)
		if target < array[mid] {
			endIndex = mid
		} else if target > mid {
			startIndex = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
