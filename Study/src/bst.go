package main

import (
	"fmt"
)

type BST struct {
	Left *BST
	Right *BST
	Value int
}

func inorder(node *BST) {
	if node == nil {
		return
	}

	inorder(node.Left)
	fmt.Println(node.Value, " ")
	inorder(node.Right)
}

func create(n int) *BST {
	var T *BST
	T = insert(T, 12)
	T = insert(T, 0)
	T = insert(T, 9)
	T = insert(T, 16)
	T = insert(T, 20)
	T = insert(T, 15)
	T = insert(T, 8)
	return T
}

func insert(T *BST, val int) *BST {

	if T == nil {
		return &BST{nil, nil, val}
	}

	if val ==  T.Value {
		return T
	}

	if val < T.Value {
		T.Left = insert(T.Left, val)
		return T
	}
	T.Right = insert(T.Right, val)
	return T
}

func main() {
	tree := create(5)
	inorder(tree)
	fmt.Println()
}
