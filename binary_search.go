package main

import (
"fmt"
)

func main(){
    mas:=[]int{2,5,7,14,23,26,27,30,34,42}
    var ch int
    fmt.Println("Введите загаданное число из списка:")
    fmt.Scan(&ch)
    binarySearch(mas,ch)
}

func binarySearch(mas []int,item int){
    low:=0
    high:=len(mas)-1
    for low<=high{
        mid:=(low+high)/2
        guess:=mas[mid]
        if guess==item{
            fmt.Printf("Позиция загаданного числа в массиве: %d",mid)
            return
        }else if guess>item{
            high=mid-1
        }else if guess<item{
            low=mid+1
        }
    }
    fmt.Println("Числа нет в списке")
}