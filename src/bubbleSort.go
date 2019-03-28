package main

func (a arr) bubbleSort() {
	for i := 0; i < len(a); i++ {
		for j := len(a) - 1; j > i; j-- {
			if a[j] < a[j-1] {
				swap(&a[j], &a[j-1])
			}
		}
	}
}
