package insertionSort

import (
//     "fmt"
)

func insertionSort(mas []int)[]int{
    for i:=1;i<len(mas);i++{
        j:=i-1
        for j>=0 && (mas[j]>mas[j+1]){
            mas[j],mas[j+1]=mas[j+1],mas[j]
            j--
        }
    }

    return mas
}