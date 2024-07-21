package greedy

//left and right -> max of both in final
//TC - O(3N)
//SC - O(2N)
func Candy_Brute(ratings []int) int {
    n:=len(ratings)

	left:=make([]int,n)
	right:=make([]int,n)

	left[0]=1
	right[n-1]=1

	//left traversal
	for i:=1;i<n;i++{
		if ratings[i]>ratings[i-1]{
			left[i] = left[i-1]+1
		}else{
			left[i]=1
		}
	}
    ans:=0
	//right traversal
	for i:=n-2;i>=0;i--{
		if ratings[i]>ratings[i+1]{
			right[i] = right[i+1]+1
		}else{
			right[i] = 1
		}
	}

 
	return ans
}

//TC - O(3N)
//SC - O(N)
func Candy_SpaceOptimize1(ratings []int)int  {
	n:=len(ratings)
    
	candies:=make([]int,n)

	for i:=range candies{
		candies[i] = 1
	}

	//left traversal
	for i:=1;i<n;i++{
		if ratings[i]>ratings[i-1]{
			candies[i] = candies[i-1]+1
		}
	}
    
	sum:= candies[n-1]

	//right traversal 
	for i:=n-2;i>=0;i--{
		if ratings[i]>ratings[i+1]{
			candies[i] = candies[i+1]+1
		}
		sum+=candies[i]
	}

	return sum
}

//TC - O(2N)
//SC - O(N)
func Candy_SpaceOptimize2(ratings []int)int  {
	n:=len(ratings)
    
	candies:=make([]int,n)
   
	candies[0]=1


	//left traversal
	for i:=1;i<n;i++{
		if ratings[i]>ratings[i-1]{
			candies[i] = candies[i-1]+1
		}else{
			candies[i]=1
		}
	}
    
	sum:=  max(1,candies[n-1])
	cur:=1

	//right traversal 
	for i:=n-2;i>=0;i--{
		if ratings[i]>ratings[i+1]{
			 cur++
		}else{
			cur=1
		}
		sum+=max(candies[i],cur)
		candies[i] = max(candies[i],cur)
	}

	return sum
}


//Slope concept
//solve for curves always
//O(N)
func Candies(ratings []int)int{
   n:= len(ratings)

   i:=1
   sum:=1
   

   for i<n{
	  //flat
       if ratings[i]==ratings[i-1]{
		   sum+=1
		   i++
		   continue
	   }

	  //up
	  peak:=1
	  for i<n && ratings[i]>ratings[i-1]{
          peak++
		  sum+=peak
		  i++
	  }

	  //down
	  down:=0
      for i<n && ratings[i]<ratings[i-1]{
         down++
		 sum+=down
		 i++
	  }

	  if down>peak{
		sum+= down-peak
	  }

   }

   return sum
}