package main

import (
	"fmt"
	"wispoz.com/structs/stack"
)

func main() {
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
