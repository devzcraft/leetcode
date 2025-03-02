package main

import "fmt"

type Node struct {
	Val       int
	Neighbors []*Node
}

func main() {
	n1 := &Node{
		Val: 1,
	}

	n2 := &Node{
		Val: 2,
	}

	n3 := &Node{
		Val: 3,
	}

	n4 := &Node{
		Val: 4,
	}
	n1.Neighbors = append(n1.Neighbors, n2, n4)
	n2.Neighbors = append(n2.Neighbors, n1, n3)
	n3.Neighbors = append(n3.Neighbors, n2, n4)
	n4.Neighbors = append(n4.Neighbors, n1, n3)

	nn := cloneGraph(n1)
	fmt.Println(nn)
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	nm := make(map[int]*Node)
	cloneRec(node, nm)

	return nm[node.Val]
}

func cloneRec(node *Node, nm map[int]*Node) *Node {
	if node == nil {
		return nil
	}

	if _, ok := nm[node.Val]; !ok {
		newNode := &Node{
			Val: node.Val,
		}

		nm[node.Val] = newNode

		for _, n := range node.Neighbors {
			newNode.Neighbors = append(newNode.Neighbors, cloneRec(n, nm))
		}
	}

	return nm[node.Val]
}
