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
	//fmt.Printf("p=%d\n",p)
	for j := low; j < high; j++ {
	    //fmt.Printf("j=%d, mas[j]=%d\n",j,mas[j])
		if mas[j] < p {
// 		    fmt.Println("")
// 		    fmt.Printf("Before: mas[%d]=%d, mas[%d]=%d\n",j,mas[j],low,mas[low])
			mas[j], mas[low] = mas[low], mas[j]
// 			fmt.Printf("After: mas[%d]=%d, mas[%d]=%d\n",j,mas[j],low,mas[low])
// 			fmt.Println(mas)
			low++
// 			fmt.Printf("low=%d\n",low)
		}
	}
//     fmt.Printf("Finish before: mas[%d]=%d, mas[%d]=%d\n",high,mas[high],low,mas[low])
	mas[low], mas[high] = mas[high], mas[low]
// 	fmt.Printf("Finish after: mas[%d]=%d, mas[%d]=%d\n",high,mas[high],low,mas[low])
// 	fmt.Println(mas)
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

