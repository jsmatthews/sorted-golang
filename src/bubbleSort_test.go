package main

import "testing"

func TestBubbleSort(t *testing.T) {
	myArr.bubbleSort()

	sortedArray := arr{0, 3, 23, 42, 50, 56, 123, 344, 456, 789}

	for i, v := range myArr {
		if v != sortedArray[i] {
			t.Error("Error: Not a sorted array")
		}
	}
}
