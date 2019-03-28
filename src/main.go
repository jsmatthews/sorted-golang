package main

import (
	"fmt"
)

type arr []int

var myArr = arr{123, 456, 789, 0, 50, 344, 23, 56, 42, 3}

func main() {
	// myArr.reset()
	// myArr.bubbleSort()

	// myArr.reset()
	// myArr.selectionSort()

	// myArr.reset()
	// myArr.insertionSort()

	myArr.reset()
	myArr.quickSort()

	fmt.Println(myArr)
}

func swap(left *int, right *int) {
	*left, *right = *right, *left
}

func (a *arr) reset() {
	*a = arr{123, 456, 789, 0, 50, 344, 23, 56, 42, 3}
}
