package TreeNode

import "fmt"

type Node struct {
	Value       string
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Println(node.Value)
}

func (node Node) SetValue() {
	node.Value = node.Value + "Changed"
	//node.print()
}
