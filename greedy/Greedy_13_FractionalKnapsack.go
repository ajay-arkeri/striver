package greedy

import (
	"sort"
)

//  values:=[]int{100,60,120}
// 	weight:= []int{20,10,30}
// 	W := 50

//TC -> O(NlogN + O(N))
//SC -> O(N)

type weightValuePair struct {
	weight int
	value  int
}

func FractionalKnapSack(values []int, weight []int, W int) float64 {

	arr := []weightValuePair{}

	for i := 0; i < len(values); i++ {
		arr = append(arr, weightValuePair{value: values[i], weight: weight[i]})
	}

	sort.Slice(arr,func(i, j int) bool {
		return (float64(arr[i].value)/float64(arr[i].weight)) > (float64(arr[j].value)/float64(arr[j].weight))
	})
    
	totalValue:=0.0

	for i:=0;i<len(values);i++{
		if arr[i].weight<W{
			totalValue+= float64(arr[i].value)
			W-=arr[i].weight
		}else{
			fraction := float64(W)/float64(arr[i].weight)
			totalValue+= fraction * float64(arr[i].value)
			break 
		}
	}
    
	return totalValue
}
