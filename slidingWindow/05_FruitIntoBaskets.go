package slidingwindow

//[1,2,3,2,2] ans=4
//TC - O(N^2)
//SC - O(1)

func TotalFruit_Brute(fruits []int) int {
	n := len(fruits)
	maxCount := 0

	for i := 0; i < n; i++ {
		mp := make(map[int]bool)
		for j := i; j < n; j++ {
			mp[fruits[j]] = true

			if len(mp) > 2 {
				break
			}

			maxCount = max(maxCount, j-i+1)
		}
	}

	return maxCount
}

//TC - O(N+N)
//SC - O(3)
func TotalFruit_Better(fruits []int) int {
	n := len(fruits)
	maxCount := 0

	mp:=make(map[int]int)

	l,r:=0,0


	for r<n{
	   mp[fruits[r]]++
	   //shrink
	   for len(mp)>2{
		  mp[fruits[l]]--  
		  if (mp[fruits[l]]==0){
			delete(mp,fruits[l])
		  }
		  l++
	   }
	   maxCount = max(maxCount,r-l+1)
	   r++
	}

	return maxCount
}


func TotalFruit_Optimal(fruits []int) int {
	n := len(fruits)
	maxCount := 0

	mp:=make(map[int]int)

	l,r:=0,0


	for r<n{
	   mp[fruits[r]]++
	   
	   if len(mp)>2{
		  mp[fruits[l]]--
		  if mp[fruits[l]]==0{
			delete(mp,fruits[l])
		  }
		  l++
	   }

	   if len(mp)<=2{
		maxCount = max(maxCount,r-l+1)
	   }

	   r++
	}

	return maxCount
}
