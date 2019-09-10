package main

import (
	TreeNode "Tree"
	"fmt"
)

func main() {

	var node = TreeNode.Node{
		Value: "test1",
		Right: &TreeNode.Node{"righttest", nil, nil},
		Left:  &TreeNode.Node{"lefttest", nil, nil},
	}

	var node2 TreeNode.Node
	node2.Value = "test2"
	node2.Left = &TreeNode.Node{"lefttest2", nil, nil}
	node2.Right = &TreeNode.Node{"righttest2", nil, nil}

	nodeArray := []TreeNode.Node{
		{Value: "test3"},
		{},
		{"test4", nil, nil},
	}

	fmt.Println(nodeArray)

	fmt.Println("print node.value as bellow")
	node.Print()

	fmt.Println("print node.value after changed as bellow")
	node.SetValue()
	node.Print()
}
