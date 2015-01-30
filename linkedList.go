package list
import (
"fmt"


)

type Node struct {
	data interface {} 
	next *Node
	prev *Node
}

type LinkedList struct {
	Front  *Node
	Back   *Node
	length int
}

func (l *LinkedList) Init() {
	n := &Node{}
	n.next = nil
	n.prev = nil
	l.Front = n
	l.Back = l.Front
	l.length = 0

}

func (l *LinkedList) Pop() interface {} {
	defer func() { l.length-- }()
	if l.Front == l.Back {
		data := l.Front.data
		l.Front.data = ""
		return data
	}
	node := l.Back
	l.Back = node.prev
	l.Back.next = nil
	data := node.data
	return data
}

func (l *LinkedList) Push(url interface {}) {
	defer func() {
		l.length++
	}()
	if l.length == 0 {

		node := l.Front
		node.data = url
		return
	}
	n := new(Node)
	n.data = url
	n.next = l.Front
	l.Front = n
	n.next.prev = n

}

func (l *LinkedList) Print() {
	var node *Node
	fmt.Printf("length:%i\n", l.length)
	for node = l.Front; node != nil; node = node.next {
		fmt.Printf("%s", node.data)
		if node.next != nil {
			fmt.Printf("->")
		} else {
			fmt.Println()
		}
	}
}

func (l LinkedList) Length() int {
	return l.length

}
