package main

func (a arr) selectionSort() {
	var min int
	for i := 0; i < len(a); i++ {
		min = i
		for j := i; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		swap(&a[min], &a[i])
	}
}
