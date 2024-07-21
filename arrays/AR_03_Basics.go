package arrays

import (
	"math"
)

//Missing number in 1 to N, numbers not sorted
//O(N * N) --> Worst case
func FindMissingNumber_Brute(arr []int)int{
	n:=len(arr)

    for i:=1;i<=n;i++{
		flag:=false

		for j:=0;j<n;j++{
			if arr[j]==i{
				flag = true
			}
		}

		if !flag{
			return i
		}
	}

	return 0
}


//Better approach --> Hashing, put everything into hash, reiterate for getting the one which is not there. 
//TC --> O(N)+O(N)

//Another approach is subtract summation from n*(n+1)/2
//TC --> O(N)
func FindMissingNumbe_Better(arr []int)int{
   
	n:=len(arr)

	expectedSum:=n*(n+1)/2
    
	realSum:=0
	
	for i:=0;i<n;i++{
       realSum+=arr[i]
	}

	return expectedSum-realSum

}


//XOR 
func FindMissingNumbe_Optimal(arr []int)int{
	XOR1, XOR2:=0,0
    
	n:=len(arr)

	for i:=0;i<n;i++{
		XOR1 = XOR1 ^ arr[i]
		XOR2 = XOR2 ^ (i+1)
	}

	return XOR1*XOR2
}

//[1,1,0,1,1,1,0,1,1]
func MaxConsecutiveOnes(arr []int)int{
    n:=len(arr)
    
	maxOnes:=math.MinInt
    
	count:=0
	for i:=0;i<n;i++{
       if arr[i]==1{
		count++
		maxOnes = max(maxOnes,count)
	   }else{
		  if count>maxOnes{
			maxOnes = count
		  }
		  count = 0
	   }
	}
    
	return maxOnes
}


func max(a,b int)int  {
	if a>b{
		return a 
	}
	return b
}

//All numbers appear twice except one
func AppearsOnce(arr []int)int  {
	XOR:=0
   
	n:=len(arr)

	for i:=0;i<n;i++{
		XOR = XOR ^ arr[i]
	}

	return XOR
}