package binarySearch

import (
	"container/heap"
	"math"
)

//TC --> O(k * n) + O(n)
//SC ---> O(N-1)
func GasStationsBrute(arr []int, k int) float64 {
	n := len(arr)

	segmentCount := make([]int, n-1)

	for i := 1; i <= k; i++ {

		//find max diff segment

		maxDiff, maxDiffIndex := -1.0, -1

		for j := 0; j < n-1; j++ {
			var diff float64 = float64((arr[j+1] - arr[j])) / (float64(segmentCount[j]) + 1.0)

			if diff > maxDiff {
				maxDiff, maxDiffIndex = diff, j
			}
		}

		segmentCount[maxDiffIndex]++
	}

	//find max after the stations placed

	maxAns := -1.0

	for i := 0; i < n-1; i++ {
		var diff float64 = float64((arr[i+1] - arr[i])) / (float64(segmentCount[i]) + 1.0)

		if diff > maxAns {
			maxAns = diff
		}
	}

	return roundFloat(maxAns,6)
}


//---------------------------------------------------------------------------------------------------------------------------------//


type segments struct{
	segmentLength float64
	segmentIndex int
}

type MaxHeap []segments

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].segmentLength > h[j].segmentLength }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(segments))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//TC --> O(N*Log(N))+ O(k * Log(N))
//SC ---> O(N-1)
func GasStations_Heap(arr []int,k int)float64{
	n:=len(arr)
    
	howMany:= make([]int,n-1)

	maxHeap:= &MaxHeap{} 
    heap.Init(maxHeap)

	//Push the diff into max heap.
	for i:=0;i<n-1;i++{
	 diff:=float64((arr[i+1]-arr[i]))
		maxHeap.Push(segments{diff,i})
	}

	for i:=1;i<=k;i++{

		//Pop the max
		maxSegment := heap.Pop(maxHeap).(segments)
		
		maxSegmentIndex:= maxSegment.segmentIndex

        howMany[maxSegmentIndex]++
          
        diff:= arr[maxSegmentIndex+1]-arr[maxSegmentIndex]

		newSegment:= segments{float64(diff)/float64(howMany[maxSegmentIndex]+1.0),maxSegmentIndex}

		heap.Push(maxHeap,newSegment)
       
	}

	return roundFloat((*maxHeap)[0].segmentLength,6)
}


func roundFloat(value float64, precision int) float64 {
	multiplier := math.Pow(10, float64(precision))
	return math.Round(value * multiplier) / multiplier
}


//----------------------------------------------------------------------------------------------------------------//


func GasStations_BS(arr []int, k int) float64{
    n:=len(arr)
    
	var low float64=0
	var high float64 =-1

	//find max diff in the array
    for i:=0;i<n-1;i++{
		high = max(high,float64(arr[i+1]-arr[i]))
	}

	// fmt.Println("Low,high = ",low,high)

	checkFunction:= func(arr []int,x float64)int{
      
		count:=0
         
		for i:=0;i<len(arr)-1;i++{
            diff:= float64(arr[i+1]-arr[i])
            segmentCount:= int(diff/x)

			if (diff) == float64(segmentCount)*x{
                segmentCount--
			}
			count+= segmentCount
		}

		return count
	}
    
    var diff float64 = 1e-6

	for high-low>diff{
          mid:= low + (high-low)/2
		//   fmt.Println(mid)

		  //check the number of stations that can be fitted.
          count := checkFunction(arr,mid)
        //   fmt.Println("count =",count)
          if count>k{
			 low = mid
		  }else{
			high = mid
		  }
	}

	return high


	
}