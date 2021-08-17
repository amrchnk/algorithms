package main

import (
    "fmt"
)

func main(){
    mas:=[]int{3,2,64,13,34,25,33,7,21,23}
    fmt.Println(mergeSort(mas))
}

func mergeSort(mas []int)[]int{
    l:=len(mas)
    if l==1{
        return mas
    }

    middle:=l/2
    left:=mas[:middle]
    right:=mas[middle:]

    return merge(mergeSort(left),mergeSort(right))
}

func merge(left, right []int)[]int{
    result:=make([]int,len(left)+len(right))
    i:=0

    for (len(left)>0)&&(len(right)>0){
        if left[0]<right[0]{
            result[i]=left[0]
            left=left[1:]
        } else{
            result[i]=right[0]
            right=right[1:]
        }
        i++
    }
    for j:=0;j<len(left);j++{
        result[i]=left[j]
        i++
    }
    for j:=0;j<len(right);j++{
            result[i]=right[j]
            i++
    }
    return result
}