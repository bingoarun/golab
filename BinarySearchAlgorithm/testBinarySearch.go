package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	index := binarySearchLoop(arr, 3, 0, len(arr))
	index1 := binarySearchRecursion(arr, 3, 0, len(arr))
	fmt.Printf("Binary search using loop:%d\n", int(index))
	fmt.Printf("Binary search using recursion:%d\n", int(index1))
}
