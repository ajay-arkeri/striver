package greedy

//Leetcode 20
// "()[]{}"
func IsValid_Easy(s string) bool {
    if len(s)==0 || len(s)%2==1{
		return false
	}

	pairs:= map[rune]rune{
       '(': ')',
	   '{' :'}',
	   '[':']',
	}

	stack:=[]rune{}
    
	for _,r:=range s{
	    if _,ok:=pairs[r];ok{
			stack = append(stack, r)
		}else if len(stack)==0 || pairs[stack[len(stack)-1]]!=r{
			return false
		}else{
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack)==0
}

//"((()))"
func IsValid_Easy2(s string)bool{
   if len(s)==0 || len(s)%2==1{
	 return false
   }

   count:=0
   
   for _,r:=range s{
	  if r=='('{
		count++
	  }else{
		count--
	  }

	  if count<0{
		return false
	  }
   }

   return count==0
}


//leetcode - 678
//TC - 3^n - worst case all *
//SC - O(n)
func CheckValidString_Brute(s string) bool {
	return recursiveHelperValidString(s,0,0)
}

func recursiveHelperValidString(s string, index int, count int )bool  {
	if count<0 {
		return false
	}

	if index==len(s){
		return count==0
	}

	char := s[index]

	if char=='('{
		return recursiveHelperValidString(s,index+1,count+1)
	}else if char==')'{
		return recursiveHelperValidString(s,index+1,count-1)
	}else{
		return recursiveHelperValidString(s,index+1,count+1) || recursiveHelperValidString(s,index+1,count-1) || recursiveHelperValidString(s,index+1,count)
	}
}


//memoized 
//O(N^2)
//O(N^2)
type state struct{
	index int
	count int
}

func CheckValidString_Memoized(s string) bool {
	memo:= make(map[state]bool)
	return recursiveHelperValidString_Memoized(s,0,0,memo)
}

func recursiveHelperValidString_Memoized(s string, index int, count int,memo map[state]bool )bool  {
    state:= state{index: index,count: count}
	if val,ok:=memo[state];ok{
		return val
	}

	if count<0 {
		return false
	}

	if index==len(s){
		return count==0
	}

	char := s[index]
    var result bool

	if char=='('{
		result = recursiveHelperValidString_Memoized(s,index+1,count+1,memo)
	}else if char==')'{
		result=  recursiveHelperValidString_Memoized(s,index+1,count-1,memo)
	}else{
		result=  recursiveHelperValidString_Memoized(s,index+1,count+1,memo) || recursiveHelperValidString_Memoized(s,index+1,count-1,memo) || recursiveHelperValidString_Memoized(s,index+1,count,memo)
	}

	memo[state]=result

	return result
}


//range based

func CheckValidString_Optimal(s string)bool{
	min,max:=0,0

	for _,r:=range s{
		if r=='('{
			min++
			max++
		}else if r ==')'{
			min--
			max--
		}else{
			min--
			max++
		}

		if min<0{
			min = 0
		}

		if max<0{
			return false
		}
	}

	return min==0
}