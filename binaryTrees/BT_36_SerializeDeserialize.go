package binarytrees

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (codec *Codec) Serialize(root *TreeNode) string {
	ans := ""

	if root == nil {
		return ans
	}

	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) > 0 {
		temp := queue[0]
		queue = queue[1:]

		if temp==nil{
			ans+="#,"
		}else{
			ans+= strconv.Itoa(temp.Val) + ","
		}
        
		if temp!=nil{
			queue = append(queue, temp.Left)
			queue = append(queue, temp.Right)
		}

	}
	return ans
}

// Deserializes your encoded data to tree.
func (codec *Codec) Deserialize(data string) *TreeNode {
     if data==""{
		return nil
	 }

	 nodes:= strings.Split(data, ",")
     rootVal,_:= strconv.Atoi(nodes[0])
	 root:= &TreeNode{Val: rootVal}
     
     queue:=[]*TreeNode{}

	 queue = append(queue, root)

	 index:=1
      
	 for len(queue)>0{
         temp:=queue[0]
         queue = queue[1:]
 
		 if nodes[index]!="#"{
			val,_:= strconv.Atoi(nodes[index])
            temp.Left = &TreeNode{Val: val}
			queue = append(queue, temp.Left)
		 }
         index++

		 if nodes[index]!="#"{
			val,_:=strconv.Atoi(nodes[index])
			temp.Right = &TreeNode{Val: val}
			queue = append(queue, temp.Right)
		 }
		 index++
	 }
    return root
}
