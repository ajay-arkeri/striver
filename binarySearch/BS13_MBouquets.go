package binarySearch

import (
	"math"
)

//m bouquets and k consequitive
//[1,10,3,10,2] m=3 k=1
func HowManyBouquets_Brute(bloomDay []int,m,k int)int{
   //If bouquets not possible
   n:= len(bloomDay)
   if m*k>n{
	return -1
   }

   //Bouquet possible
   min,max:= findMinMax(bloomDay)
   
   for i:=min;i<=max;i++{
	  if possible(bloomDay,m,k,i){
		return i
	  }
   }
   
   return -1
}
//TC --> O(n * (max-min))

func findMinMax(arr []int)(int,int){
	min,max:= math.MaxInt, math.MinInt
   
	n:=len(arr)

	for i:=0;i<n;i++{
      
		if arr[i]>max{
			max = arr[i]
		}

		if arr[i]<min{
			min = arr[i]
		}

	}
	return min,max
}

func possible(arr []int, m,k,n int)bool{
    count:=0
    bouquet:=0
    
    for i:=0;i<len(arr);i++{
        if arr[i]<=n{
            count++
        }else{
            bouquet += count/k
            count = 0
        }
    }
    bouquet+=count/k
    
    return bouquet>=m
}

func HowManyBouquets(bloomDay []int,m,k int)int{
   //If bouquets not possible
   n:= len(bloomDay)
   if m*k>n{
	return -1
   }

   //Bouquet possible
   min,max:= findMinMax(bloomDay)
   
   ans:= max
 
    for min<=max{
        mid:= min+(max-min)/2

        if possible(bloomDay,m,k,mid){
            if mid<ans{
                ans = mid
            }
            max = mid-1
        }else{
            min = mid+1
        }
    }

    return ans
   
}

//TC --> O(n * log(max-min))