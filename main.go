package main

import (
	"fmt"
	slidingwindow "striver/slidingWindow"
)

func main() {
	//  node:= &binarytrees.TreeNode{Val: 9}

	//  arr := []int{12,5,1,8,14,11,6,13,15}

	// // arr:= []int{1,3}
	//  for _,Val:= range arr{
	// 	node.Insert(Val)
	//  }

	//fmt.Println(binarytrees.CountNode_Brute(node))

	fmt.Println(slidingwindow.MinWindow_Optimal("ADOBECODEBANC","ABC"))
	fmt.Println(slidingwindow.MinWindow_Optimal("aa","aa"))
	
}

