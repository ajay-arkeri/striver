package arrays

import (
	"math"
	"sort"
)

//{102,4,100,1,101,3,2,1,}
//TC -- O(N^2)
func LongestConsecutiveSequence_Brute(nums []int) int {
	ans := math.MinInt
	n:=len(nums)

	count:=1

	for i:=0;i<n;i++{
		x:=nums[i] + 1

		for {
			flag:=false
			for j:=0;j<n;j++{
				if nums[j]==x{
                   x++
				   count++
	               ans = max(ans,count)
				   flag = true
				}
			}
			if !flag{
				count = 1
				break	
			}
		}
	}

    return ans
}

//Sorting 
func LongestConsecutiveSequence_Better(nums []int)int{
   n:=len(nums)
   
   if n==0{
	return 0
   }

   sort.Ints(nums)


   prev:=math.MinInt
   count:=0
   ans:=1

   for i:=0;i<n;i++{
	 if nums[i] == prev + 1{
        count++
        prev = nums[i]
	 }else if nums[i]!=prev{
	    count = 1 
		prev = nums[i]
	 }

	 ans = max(ans,count)
   }

   return ans
}

func LongestConsecutiveSequence_Optimal(nums []int)int{
	n:=len(nums)
    
    mp:=make(map[int]bool)

	for i:=0;i<n;i++{
        mp[nums[i]] = true
	}
    
	ans:=0

	for num:= range mp{
		if mp[num-1]{
			continue
		}

		count:=1

		for mp[num+1]{
            count++
			num++
		}

		ans = max(ans,count)
	}

	return ans	
}