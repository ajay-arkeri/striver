package arrays

import "fmt"

//TC - O(N)+O(N/2)
//SC - O(N)
func RearrangeArray_Brute(nums []int) []int {

	arr1, arr2 := []int{}, []int{}
	n := len(nums)

	for i := 0; i < n; i++ {
		if nums[i] < 0 {
			arr2 = append(arr2, nums[i])
		} else {
			arr1 = append(arr1, nums[i])
		}
	}

	for i := 0; i < n/2; i++ {
		nums[2*i] = arr1[i]
		nums[2*i+1] = arr2[i]
	}

	return nums
}

//TC --> O(N)
//SC ---> O(N)
func RearrangeArray(nums []int) []int {
    n:=len(nums)
    
	result:=make([]int,n)

    positiveIndex,negativeIndex:=0,1

    for i:=0;i<n;i++{
		if nums[i]<0{
			result[negativeIndex] = nums[i]
			negativeIndex+=2
		}else{
			result[positiveIndex] = nums[i]
			positiveIndex+=2
		}
	}

	return result
}


//When positives and negatives are not equal

func RearrangeArray_Type2(nums []int) []int {
    n:=len(nums)
    
	postives,negatives:=[]int{},[]int{}
    resultant:=make([]int,n)

	for i:=0;i<n;i++{
		if nums[i]>0{
			postives = append(postives, nums[i])
		}else{
			negatives = append(negatives, nums[i])
		}
	}

	fmt.Println(postives,negatives)

	postivesLen,negativesLen := len(postives),len(negatives)


	if postivesLen>negativesLen{
        for i:=0;i<negativesLen;i++{
            resultant[2*i]=postives[i]
			resultant[2*i+1] = negatives[i]
		}
        
		index:=negativesLen*2
		for i:=negativesLen;i<postivesLen;i++{
			resultant[index] = postives[i]
			index++
		}

	}else{
           for i:=0;i<postivesLen;i++{
            resultant[2*i]=postives[i]
			resultant[2*i+1] = negatives[i]
		}

		for i:=postivesLen;i<negativesLen;i++{
			resultant[i] = negatives[i]
		}
	}

	return resultant
}