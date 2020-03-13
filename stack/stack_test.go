package stack

import "testing"

func TestStack_Push(t *testing.T) {
	Stack := Stack{}
	Stack.Push(1)
	val, err := Stack.Pop()
	if val != 1 {
		t.Fatal("Test Fail, must be 1")
	}
	if err == false {
		t.Fatal("Test Fail, must be 1")
	}
}

func TestStack_Pop(t *testing.T) {
	Stack := Stack{}
	Stack.Push(1)
	val, err := Stack.Pop()
	if val != 1 {
		t.Fatal("Test Fail, must be 1")
	}
	if err == false {
		t.Fatal("Test Fail, must be 1")
	}
}

func TestStack_Peek(t *testing.T) {
	Stack := Stack{}
	Stack.Push(1)
	root, err := Stack.Peek()
	if err == false {
		t.Fatal("root not set")
	}
	if root.value != 1 {
		t.Fatal("Value of root must be 1")
	}
}

func TestStack_Size(t *testing.T) {
	Stack := Stack{}
	Stack.Push(1)
	if Stack.Size() != 1 {
		t.Fatal("Size must be 1")
	}
}

func TestStack_Peek2(t *testing.T) {
	Stack := Stack{}

	node, err := Stack.Peek()
	if err != false {
		t.Fatal("Err must be true")
	}
	if node.value != 0 {
		t.Fatal("Value must be 0")
	}

}
func TestStack_Pop2(t *testing.T) {
	Stack := Stack{}

	value, err := Stack.Pop()
	if err != false {
		t.Fatal("Err must be true")
	}
	if value != 0 {
		t.Fatal("Value must be 0")
	}

}
