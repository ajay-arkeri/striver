package greedy

import (
	"math"
	"sort"
)

//  arr:=[]int{900,940,950,1100,1500,1800}
//  dep:= []int{910,1200,1120,1130,1900,2000}

func MinimumPlatforms_Brute(arr []int, dep []int) int {
	n := len(arr)

	minimumPlatforms := math.MinInt
    
	for i:=0;i<n;i++{
		count:=1
        
		for j:=0;j<n;j++{
			if i==j{
				continue
			}else{
				if arr[i]>=arr[j] && dep[j]>=arr[i]{
                     count++
				}
			}
		}

		minimumPlatforms = max(minimumPlatforms,count)
	}
   
	return minimumPlatforms
}

type pair struct{
	time int
	arrival bool
}

//TC - O(2nlogn) + o(n)
//SC - O(2n)
func MinimumPlatforms_Better(arr []int, dep []int)int{
	combined := []pair{}
    
	miniumPlatform:= math.MinInt

   	for i:=0;i<len(arr);i++{
		//arrival
		combined = append(combined, pair{time: arr[i] , arrival: true})

		//departure
		combined = append(combined, pair{time: dep[i], arrival: false})
	}

	sort.Slice(combined,func(i, j int) bool {
		return combined[i].time<combined[j].time
	})
    
	count:=0
	for _,val:= range combined{
         if val.arrival{
            count++
		 }else{
            count--
		 }

		 miniumPlatform =max(miniumPlatform,count)
	}

	return miniumPlatform
}

//TC- O(2nlogn) + o(2n)
//SC - O(1)
func MinimumPlatforms_Optimal(arr []int,dep []int)int  {
	minimumPlatforms:= math.MinInt
    
	sort.Ints(arr)
	sort.Ints(dep)

	i:=0
    j:=0
  
	count:=0 

	for i<len(arr){
		if arr[i]<dep[j]{
            count++
			i++
		}else{
			count--
			j++
		}

		minimumPlatforms = max(minimumPlatforms,count)
	}

	return minimumPlatforms
}