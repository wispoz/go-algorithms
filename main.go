package main

import (
	"fmt"
	"wispoz.com/structs/queue"
	"wispoz.com/structs/stack"
)

func main() {

	q := queue.Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	fmt.Println(q.Size())

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())

	fmt.Println(q.Size())

	Stack := stack.Stack{}

	Stack.Push(1)
	Stack.Push(2)
	Stack.Push(3)

	fmt.Println(Stack.Size())

	fmt.Println(Stack.Pop())
	fmt.Println(Stack.Pop())
	fmt.Println(Stack.Pop())

	fmt.Println("stack.size")
	fmt.Println(Stack.Size())
	fmt.Println("done")
}
