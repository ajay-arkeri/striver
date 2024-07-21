package miscellaneous

import "math"



func CoinChange(coins []int, amount int) int {
    memo:= make(map[int]int)
	ans:= allCoinCombinations(coins,amount,memo)

    if ans==math.MaxInt32{
		return -1
	}

	return ans	
}

func allCoinCombinations(coins []int, amount int,memo map[int]int)int{
   
	if amount==0{
       return 0
	}

	if amount<0{
		return math.MaxInt32
	}

	if val,ok:=memo[amount];ok{
		return val
	}

    minCount:=math.MaxInt32
	for i:=0;i<len(coins);i++{
		count:= 1 +  allCoinCombinations(coins,amount-coins[i],memo)

        minCount = min(minCount,count)
	}
   
	memo[amount] = minCount

	return minCount 
	
}


