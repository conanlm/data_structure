package main

import (
	"fmt"
	"os"
)

type Node struct {
	num  int
	next *Node
}

func createList(size int) *Node {
	if size < 0 {
		os.Exit(-1)
	}
	node := new(Node)
	node.num = 1
	q := node
	for i := 2; i <= size; i++ {
		list := new(Node)
		list.num = i
		node.next = list
		node = list
	}
	node.next = q
	return node
}

func empty(list *Node) bool {
	if list.next == list {
		return true
	}
	return false
}

func traverse(list *Node) {
	if empty(list) {
		return
	}
	for p := list.next; p != list; p = p.next {
		fmt.Printf("%5d", p.num)
	}
	fmt.Printf("\n%d \n", list.num)
}

func length(list *Node) int {
	i := 1
	for p := list.next; p != list; p = p.next {
		i++
	}
	return i
}

func jose(list *Node, n int) {
	node := list.next
	for count := 1; 1 < length(node); count++ {
		for i := 1; i < n-1; i++ {
			node = node.next
		}
		kill := node.next
		fmt.Printf("第%d个出局的人为：%3d号\n", count, kill.num)
		node.next = kill.next
		fmt.Println(length(kill))
		node = node.next
		fmt.Println(length(node))
		kill = nil
	}
	fmt.Println("最后获胜的是: ", node.num)
}

func main() {
	list := createList(41)
	traverse(list)
	jose(list, 3)
}
