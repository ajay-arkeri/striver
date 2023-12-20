package binarySearch

//Minimum weight of ship  with days

func ShipWithinDays_Brute(weights []int, days int) int {
    min,max:= maxInArray(weights),sumOfWeights(weights)

	for i:=min;i<=max;i++{
       if possibleToShip(weights,days,i){
		return i
	   }
	}

	return max
}

func possibleToShip(weights []int, days int, w int )bool{
     count:=0
     time:=0
	 for i:=0;i<len(weights);i++{
         if count+weights[i]>w{
			time+=1
			count = weights[i]
		 }else if count+weights[i]==w{
			time+=1
			count = 0
		 }else{
             count+= weights[i]
		 }
	 }

	 if count!=0{
		time+=1
	 }
     
	 return time<=days
}


func possibleToShip_Optimal(weights []int, days int, w int )bool{
     load:=0
     time:=1
	 for i:=0;i<len(weights);i++{
         if load+weights[i]>w{
			time+=1
			load = weights[i]
		 }else{
             load+= weights[i]
		 }
	 }
     
	 return time<=days
}


func sumOfWeights(arr []int)int  {
	sum:=0

	for i:=0;i<len(arr);i++{
		sum+= arr[i]
	}

	return sum
}


func ShipWithinDays(weights []int, days int) int {
      min,max:=maxInArray(weights),sumOfWeights(weights)
    
    ans:=max
    for min<=max{
        mid:= min+(max-min)/2

        if possibleToShip_Optimal(weights,days,mid){
            if mid<ans{
                ans = mid
            }
            max = mid-1
        }else{
            min = mid+1
        }
    }

    return ans
}
