package binarySearch

//Book Allocation / Painters partition / Split Array largest sum   ---> min(max)

//arr = [25,46,28,49,24] m=4               
//each student must get 1 and should allocate all books
//answer lies between [max, sum] [49,172]

func BookAllocation_brute(arr []int, m int)int{
   min,max:= maxInArray(arr),sumOfArray(arr)

   for i:=min;i<=max;i++{
	  students:= calculateStudents(arr,i)
     
	  if students==m{
		return i
	  }
   }

   return min
}

func calculateStudents(arr []int,size int)int{
	students,PageCount:=1,0

	for i:=0;i<len(arr);i++{              
		if arr[i]+PageCount<=size{         
	        PageCount+=arr[i]
		}else{
			students++
			PageCount=arr[i]
		}
	}

	return students
}

func sumOfArray(arr []int)int{
	sum:=0

	for i:=0;i<len(arr);i++{
		sum+= arr[i]
	}

	return sum
}


func BookAllocation(arr []int, m int)int{
	low,high:= maxInArray(arr),sumOfArray(arr)

	for low<=high{
		mid:=low+(high-low)/2

		students:= calculateStudents(arr,mid)

		if students<=m{
            high = mid -1 
		}else{
			low = mid + 1
		}
	}

	return low
}