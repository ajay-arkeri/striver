package greedy

//[5,5,5,10,20]

//TC- O(N)
//SC - O(1)

func LemonadeChange(bills []int) bool {
    fives,tens:= 0,0

	for i:=0;i<len(bills);i++{
		if bills[i]==5{
           fives+=1
		}else if bills[i]==10{
           if fives>=1{
              tens+=1
			  fives-=1
		   }else{
			  return false
		   }
		}else{
           if fives>0 && tens>0{
			 tens-=1
			 fives-=1
		   }else if fives>=3{
			  fives-=3
		   }else{
			return false
		   }
		}
	}

	return true
}