package insertionSort

import(
    "testing"
    "math/rand"
    "fmt"
)

func RandMas(n int)[]int{
    mas:=[]int{}
    for i:=0;i<n;i++{
        mas=append(mas,rand.Intn(100))
    }
    return mas
}

func TestInsertionSort(t *testing.T){
    testS:=[]int{5,12,20,16}
    for _,item:=range testS{
        var mas []int
        mas=RandMas(item)
        fmt.Println("original array: ",mas)
        res:=insertionSort(mas)
        fmt.Println("sort array: ",res)
        fmt.Println()
    }
}