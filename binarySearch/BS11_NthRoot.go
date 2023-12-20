package binarySearch

import "fmt"

//N=3  M=27  ans = 3
//N=4  M=69  ans = -1


func NthRoot_Brute(n int, m int)int{
	ans:= -1
    
	//learn power exponential method
	power:= func(n,m int )int{
		 ans:= 1
		 for i:=1;i<=m;i++{
			ans = ans *n
		 }
		 return ans
	}
    for i:=1;i<=m;i++{
       if power(i,n) ==m{
		return i
	   }
	}

	return ans
}
//TC ---> O(M *N)
//TC --> O(M * log2N)  --> once learned


func NthRoot(n, m int)int  {
    
	//mid^n
    power:= func(n,m int)int{
		ans:= 1
		for i:=1;i<=n;i++{
           ans = ans * m
		}
        fmt.Println(ans)
		return ans
	}

	low,high:= 1,m

	for(low<=high){
		mid:= low+(high-low)/2

		temp := power(n,mid)

		if temp==m{
			return mid
		}

		if temp>m{
			high = mid-1
		}else{
			low = mid+1
		}
	}

	return -1
}

//TC --> O(log2M * N)

//Overflow case   n=10 m=10^9   mid=10^90  ---> So rewrite the power function
// (n,mid) --> 


func NthRoot_Optimal(n int, m int)int{
   low,high:= 1,m

	for(low<=high){
		mid:= low+(high-low)/2
       // fmt.Print(mid,"=")
		temp := powerOptimized(n,mid,m)
       // fmt.Println(temp)
		if temp==0{
			return mid
		}

		if temp==1{
			low = mid+1
		}else{
			high = mid-1
		}
	}
	return -1
}


//if ans>m return 2
//if ans ==m return 0
//if ans < m return 1

func powerOptimized(n int, mid int, num int)int{
   ans:=1

   for i:=1;i<=n;i++{
      ans = ans * mid

	  if ans>num{
		return 2
	  }
   }

   if ans==num{
	return 0
   }
   return 1
}