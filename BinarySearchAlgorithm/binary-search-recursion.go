/*
	Binary search implementation in Go using recursion
*/
package main

func binarySearchRecursion(array []int, target int, lowIndex int, highIndex int) int {
	if lowIndex > highIndex {
		return -1
	}
	mid := int((lowIndex + highIndex) / 2)
	if target < array[mid] {
		return binarySearchRecursion(array, target, lowIndex, mid)
	} else if target > array[mid] {
		return binarySearchRecursion(array, target, mid+1, highIndex)
	} else {
		return mid
	}
}
