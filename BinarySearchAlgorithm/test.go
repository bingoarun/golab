package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	index := binarySearchLoop(arr, 3, 0, len(arr))
	fmt.Println(string(index))
}
