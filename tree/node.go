package tree

import (
	"fmt"
)

type Node struct {
	Value       int
	Left, Right *Node
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}

func (node Node) Print() { //在方法名前面加上一个括号，叫做接受者
	fmt.Println(node.Value)
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("setting Value to nil node.Ingore")
		return
	}
	node.Value = value

}
