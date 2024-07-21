package arrays

import (
	"math"
)

//TC-O(N^2)
func MaxSubArray_Brute(nums []int) int {
    n:=len(nums)
	maxSum:=math.MinInt
 
	for i:=0;i<n;i++{
		sum:=0
		for j:=i;j<n;j++{
           sum+=nums[j]

		   if sum>maxSum{
			maxSum = sum
		   }
		}
	}

	return maxSum	
}

//Kadane Algo --> if sum<0 neglect the subarray
func MaxSubArray_Optimal(nums []int) int {
    n:=len(nums)
	maxSum:=math.MinInt
    
	sum:=0

	for i:=0;i<n;i++{
       sum+=nums[i]

	   if sum>maxSum{
		maxSum = sum
	   }

	   if sum<0{
		sum = 0
	   }

	}

	return maxSum	
}


//Return i,j of subarray
func MaxSubArray_Optimal2(nums []int) []int {
    n:=len(nums)
	maxSum,ansStart,ansEnd,start:=math.MinInt,-1,-1,-1
	sum:=0

	for i:=0;i<n;i++{
       if sum==0{
		start = i
	   }
       sum+=nums[i]

	   if sum>maxSum{
		maxSum = sum
		ansStart = start
		ansEnd = i
	   }

	   if sum<0{
		sum = 0
	   }

	}

	return []int{ansStart,ansEnd}	
}


//Best time to buy and sell a stock
func MaxProfit_Brute(prices []int) int {
        n:=len(prices)
    
    maxProfit:=0
    profit:=0

    for i:=0;i<n-1;i++{
        for j:=i+1;j<n;j++{
             profit = prices[j]-prices[i]

             if profit>maxProfit{
                 maxProfit = profit
             }
        }
    }

    return maxProfit
}

func MaxProfit_Optimal(prices []int)int{
	  n:=len(prices)
    
    maxProfit,profit:=0,0
    min:=math.MaxInt

    for i:=0;i<n;i++{
       if prices[i]<min{
           min = prices[i]
       }
      
       profit = prices[i]-min
	
       if profit>maxProfit{
           maxProfit = profit
       }
    }

    return maxProfit
}