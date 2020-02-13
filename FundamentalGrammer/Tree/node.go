package tree

import (
	"fmt"
)

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Cannot set value for nil treeNode pointer")
		return
	}
	node.Value = value
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}
