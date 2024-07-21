package binarytrees

import (
	"container/heap"
	"fmt"
	"sort"
)

type pair struct {
	node *TreeNode
	col  int
}

func TopView_Sort(root *TreeNode) []int {
	ans := []int{}

	if root == nil {
		return ans
	}

	mp := make(map[int]int)

	queue := []pair{}

	queue = append(queue, pair{node: root, col: 0})

	for len(queue) > 0 {
		ds := queue[0]
		queue = queue[1:]

		node := ds.node
		dist := ds.col

		if _, ok := mp[dist]; !ok {
			mp[dist] = node.Val
		}

		if node.Left != nil {
			queue = append(queue, pair{node: node.Left, col: dist - 1})
		}

		if node.Right != nil {
			queue = append(queue, pair{node: node.Right, col: dist + 1})
		}
	}
    
	keys:=[]int{}

	for key:= range mp{
        keys = append(keys, key)
	}
  
	sort.Ints(keys)

	for _,key:= range keys{
		ans = append(ans, mp[key])
		fmt.Println(key,mp[key])
	}

	return ans
}

type minHeap []pair

func (h minHeap)Len()int{return len(h)}
func (h minHeap)Less(i,j int)bool {return h[i].col<h[j].col}
func (h minHeap)Swap(i,j int){h[i],h[j]= h[j],h[i]}

func (h *minHeap)Push(x interface{}){
   *h = append(*h, x.(pair))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func TopView_Heap(root *TreeNode)[]int  {
	ans:=[]int{}

	if root==nil{
		return ans
	}
    
	minHeap:= &minHeap{}
    mp:=make(map[int]int)
	queue:= []pair{}
    queue = append(queue, pair{node: root,col: 0})
    
	for len(queue)>0{
        temp:=queue[0]
        queue = queue[1:]
 
		node:=temp.node
        dist := temp.col
        
        if _,ok:=mp[dist];!ok{
			mp[dist] = node.Val
            heap.Push(minHeap,temp)
		}       

        if node.Left!=nil{
			queue = append(queue, pair{node: node.Left,col: dist-1})
		}

		if node.Right!=nil{
			queue = append(queue, pair{node: node.Right,col: dist+1})
		}
	}

	for len(*minHeap)>0{
		item:= heap.Pop(minHeap).(pair)

		ans = append(ans, item.node.Val)
	}

	return ans
}