package miscellaneous

//739
func DailyTemperatures_Brute(temperatures []int) []int {
	n := len(temperatures)

	ans := make([]int, n)

	for i := 0; i < n; i++ {
		flag := false
		count := 0
		for j := i + 1; j < n; j++ {
			count++
			if temperatures[j] > temperatures[i] {
				ans[i] = count
				flag = true
				break
			}
		}

		if !flag {
			ans[i] = 0
		}
	}

	return ans
}

//[73,74,75,71,69,72,76,73]
func DailyTemperatures(temperatures []int) []int {
     n:=len(temperatures)
     ans:=make([]int,n)
	 stack:=make([]int,n)

	 for i:=n-1;i>=0;i--{
		
  		for len(stack)>0 && temperatures[i]>temperatures[stack[len(stack)-1]]{
            stack = stack[:len(stack)-1]
		}
        
		if len(stack)>0{
           ans[i] =  stack[len(stack)-1] -i
		}else{
			ans[i] = -1
		}

		stack = append(stack, i)
        
	 }

	 return ans
}