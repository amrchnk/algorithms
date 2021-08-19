package binarySearch

import (
    "sort"
)

func binarySearch(mas []int,item int)int{
    sort.Ints(mas)
    low:=0
    high:=len(mas)-1
    for low<=high{
        mid:=(low+high)/2
        guess:=mas[mid]

        if guess==item{
            return mid
        }else if guess>item{
            high=mid-1
        }else if guess<item{
            low=mid+1
        }
    }
    return -1
}