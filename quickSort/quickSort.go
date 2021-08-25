package main

import (
	"fmt"
)
func main() {
	list := []int{12,5,43,17,6,8,34,9,22,20}

	quickSort(list, 0, len(list)-1)
	fmt.Println(list)
}
func partition(mas []int, low, high int) int {
	p := mas[high]
	for j := low; j < high; j++ {
		if mas[j] < p {
			mas[j], mas[low] = mas[low], mas[j]
			low++
		}
	}
	mas[low], mas[high] = mas[high], mas[low]
	return low
}

func quickSort(mas []int, low, high int) {
	if low > high {
		return
	}

	p := partition(mas, low, high)
	quickSort(mas, low, p-1)
	quickSort(mas, p+1, high)
}

