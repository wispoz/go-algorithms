package stack

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	root *Node
	size int
}

func (stack *Stack) Pop() (int, bool) {
	if stack.size == 0 {
		return 0, false
	}
	oldRoot := stack.root
	stack.root = oldRoot.next
	stack.size--
	return oldRoot.value, true
}

func (stack *Stack) Push(value int) {
	oldRoot := stack.root
	stack.root = &Node{value: value, next: oldRoot}
	stack.size++
}

func (stack *Stack) Size() int{
	return stack.size
}

func (stack *Stack) Peek() (*Node, bool) {
	if stack.size == 0 {
		return &Node{}, false
	}
	return stack.root, true
}

func main() {
	fmt.Println("test")
}
