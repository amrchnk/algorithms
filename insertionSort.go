package main

import (
    "fmt"
)

func main(){
    mas:=[]int{2,4,3,6,34,12,64,23,22}
    fmt.Println(insertionSort(mas))
}

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