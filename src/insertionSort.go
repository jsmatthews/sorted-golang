package main

func (a arr) insertionSort() {
	for i := 1; i < len(a); i++ {
		j := i

		for j > 0 {
			if a[j-1] > a[j] {
				swap(&a[j-1], &a[j])
			}

			j--
		}
	}
}
