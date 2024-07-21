package arrays

import (
	"reflect"
	"sort"
)

//Lets do permutation the recursive way

//TC --> O(N! * N)
//SC --> O(N) + O(N) marking map

func Permute_Brute1(nums []int) [][]int {
	n:=len(nums)
	marked:=make(map[int]bool,n)
	temp:=[]int{}
	ans:= [][]int{}

	var generatePermutations func()

    generatePermutations = func(){
        //base condition
		if len(temp)==n{
		    dummy:=make([]int,n)
			copy(dummy,temp)
			ans = append(ans, dummy)
			return
		}

		for i:=0;i<n;i++{
			if !marked[i]{
				marked[i] = true
                temp = append(temp, nums[i])
				generatePermutations()
				marked[i]=false
				temp = temp[:len(temp)-1]
			}
		}
	}

	
    generatePermutations();
	return ans
}


func Permute_Brute2(nums []int) [][]int {
	n:=len(nums)
	marked:=make(map[int]bool,n)
	ans:= [][]int{}

    generatePermutations(&ans,marked,[]int{},nums,n);
	return ans
}

func generatePermutations (ans *[][]int,marked map[int]bool,temp []int,nums []int,n int){
        //base condition
		if len(temp)==n{
		    dummy:=make([]int,n)
			copy(dummy,temp)
			*ans = append(*ans, dummy)
			return	
		}

		for i:=0;i<n;i++{
			if !marked[i]{
				marked[i] = true
                temp = append(temp, nums[i])
				generatePermutations(ans,marked,temp,nums,n)
				marked[i]=false
				temp = temp[:len(temp)-1]
			}
		}
}

//TC -->O(N! * N)
//SC --> O(1)  neglecting recursion stack space depth, and ans space to return 
func Permute_Better(nums []int)[][]int{
    n:=len(nums)
    ans:=[][]int{}

    generatePermutationsSwapping(0,nums,&ans,n)

	return ans
}

func generatePermutationsSwapping(index int,nums []int,ans *[][]int, n int){
    //base condition
	if index==n{
		temp:=make([]int,n)
		copy(temp,nums)
		*ans = append(*ans, temp)
	}


	for i:=index;i<n;i++{
       swap(nums,index,i)
	   generatePermutationsSwapping(index+1,nums,ans,n)
	   swap(nums,index,i)
	}
}

func swap(nums []int,i,j int){
    nums[i],nums[j] = nums[j],nums[i]
}


// -----------------------------------------------------------------------------

//Next permuatation
//This works only for unique elements
func NextPermutation_Brute(nums []int)  {
    n:=len(nums)

    temp:=make([]int,n)

	copy(temp,nums)
 
	sort.Ints(temp)
    
	marked:= make(map[int]bool,n)
	allPermutations:= [][]int{}

	generatePermutations(&allPermutations,marked,[]int{},temp,n)

	for i:=0;i<len(allPermutations);i++{
		if reflect.DeepEqual(allPermutations[i],nums){

		    if i==len(allPermutations)-1{
                copy(nums,allPermutations[0])
			}else{
                copy(nums,allPermutations[i+1])			
			}

			return
		}
	}
}

//Optimal
//TC --> O(3N)
//SC --> O(1)
func NextPermutation_Optimal(nums []int){
    n:=len(nums)
    
	index:=-1

	//find the the lowest from the last
	for i := n-2; i >=0; i-- {
		if nums[i]<nums[i+1]{
			index = i
			break	
		}
	}

	if index==-1{
		reverse(nums,0,n-1)
		return
	}
    

	//find second lowest from last and swap
	for i:=n-1;i>=index;i--{
		if nums[i]>nums[index]{
           nums[i],nums[index] = nums[index],nums[i]
		   break	
		}
	}

	reverse(nums,index+1,n-1) 
} 

func reverse(nums []int,start,end int){
	for i,j:=start,end;i<j;i,j = i+1,j-1{
		nums[i],nums[j] = nums[j],nums[i]
	}
}
