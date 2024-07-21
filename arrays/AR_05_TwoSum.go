package arrays

import "sort"

//near O(N^2) solution
func TwoSum_Brute(nums []int, target int) []int {
    n:=len(nums)

	for i:=0;i<n;i++{
		for j:=i+1;j<n;j++{
			if nums[i]+nums[j]==target{
				return []int{i,j}
			}

			// if i!=j  --> for j:=0 j<n
		}
	}

	return []int{}
}


//Hash Map
//TC -- O(N*1) no collisions O(N*N) collisions
//SC -- O(N)
func TwoSum_Better(nums []int, target int) []int {
    n:=len(nums)
   
	mp:=make(map[int]int,n)

	for i:=0;i<n;i++{
	   if val,ok:=mp[target-nums[i]];ok{
		  return []int{val,i}
	   } 

	   mp[nums[i]] = i
	}

	return []int{}
}


//Solution best only to return true and false
func TwoSum_Optimal1(nums []int, target int) bool {
    n:=len(nums)
    
	sort.Ints(nums)
	i,j:=0,n-1

	for i<j{
        
		sum:= nums[i]+nums[j]

		if sum==target{
			return true
		}else if sum>target{
			j--
		}else{
			i++
		}
	}

	return false
}


//TC --> O(N)+O(NlogN)+O(N)
type tuple struct{
	num int
	index int
}

func TwoSum_Optimal2(nums []int, target int) []int{
    n:=len(nums)
    
	modifiedNums:=[]tuple{}

	for i:=0;i<n;i++{
       modifiedNums = append(modifiedNums, tuple{nums[i],i})
	}

	sort.Slice(modifiedNums,func(i, j int) bool {
		return modifiedNums[i].num<modifiedNums[j].num
	})

	i,j:=0,n-1


	for i<j{
		sum:=modifiedNums[i].num+modifiedNums[j].num

		if sum==target{
			return []int{modifiedNums[i].index,modifiedNums[j].index}
		}else if sum>target{
			j--
		}else{
			i++
		}
	}

	return []int{}
	
}