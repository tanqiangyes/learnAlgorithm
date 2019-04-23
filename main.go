package main

import (
	"fmt"
	"github.com/tanqiangyes/learnAlgorithm/list"
)

func main() {
	node := list.Node{}
	for i := 0; i< 10 ; i++  {
		list.InsertNode(&node,i,i)
	}
	fmt.Println(node)

	fmt.Println(list.GetNodeData(node,3))
	fmt.Println(node)

	fmt.Println(list.DeleteNode(&node,3))
	fmt.Println(node)
}
