package binarySearch

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type testPair struct {
	ItemsCount int
	ch int
}

var testData =[]testPair{
	{20,14},
	{25,7},
	{12,4},
	{34,18},
}

func GetIntMas(items int)[]int{
	var mas []int
	for i:=0;i<items;i++{
		mas=append(mas, i)
	}
	return mas
}


func TestBinarySearch(t *testing.T){
	for _,item:=range testData{
		var mas []int=GetIntMas(item.ItemsCount)
		res:=binarySearch(mas,item.ch)
		assert.Equal(t, item.ch, res)
	}
}