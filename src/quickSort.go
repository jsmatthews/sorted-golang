package main

import "math/rand"

func (a arr) quickSort() arr {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1
	pivot := rand.Int() % len(a)

	swap(&a[right], &a[pivot])

	for i := range a {
		if a[i] < a[right] {
			swap(&a[i], &a[left])
			left++
		}
	}

	swap(&a[right], &a[left])

	a[:left].quickSort()
	a[left+1:].quickSort()

	return a
}
