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
		} else if target > array[mid] {
			startIndex = mid
		} else {
			return mid
		}
	}
	return -1
}
