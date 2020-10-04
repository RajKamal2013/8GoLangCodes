package main
import  (
	"fmt"
)

type Node struct {
	Value int
	Next *Node
}

func addNode(node *Node, val int) int {
	if root == nil {
		node = &Node{val, nil}
		root = node
		return 0
	} else {
		if val == node.Value {
			fmt.Println("Node already exist ")
			return -1
		}

		if node.Next == nil {
			node.Next = &Node{val, nil}
			return -2
		}

		return addNode(node.Next, val)
	}
}

func traverseList(node *Node)  {
	if node == nil {
		fmt.Println("Empty List ")
		return;
	}  else {
		var p *Node
		p = node
		for p != nil {
			fmt.Print(p.Value, " ")
			p = p.Next
		}
	}
	fmt.Println()
}

var root = new(Node)

func main() {
	fmt.Println(root)
	root = nil
	traverseList(root)
	addNode(root, 1)
	addNode(root, 2)
	addNode(root, 3)
	traverseList(root)
}
