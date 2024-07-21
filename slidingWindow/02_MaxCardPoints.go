package slidingwindow

//Leetcode 1423
func MaxScore(cardPoints []int, k int) int {
    lSum,rSum:=0,0
  
    for i:=0;i<k;i++{
        lSum+=cardPoints[i]
    }

    maxSum:=lSum
    
    rightIndex:=len(cardPoints)-1

    for i:=k-1;i>=0;i--{
        lSum-= cardPoints[i]
        rSum +=cardPoints[rightIndex]
        rightIndex--

        maxSum = max(maxSum,lSum+rSum)
    }   
 
    return maxSum
}