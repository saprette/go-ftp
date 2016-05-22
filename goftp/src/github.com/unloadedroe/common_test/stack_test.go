package common_test

import (
	"github.com/unloadedroe/common"
	"testing"
)

func TestStackPeek(t *testing.T) {
	stack := new(common.Stack)
	for i := 3; i >= 0; i-- {
		stack.Push(i)
	}
	if stack.Peek() != 0 {
		t.Fatal("Incorrect value.")
	}
}

func TestStackPop(t *testing.T) {
	stack := new(common.Stack)
	for i := 3; i >= 0; i-- {
		stack.Push(i)
	}
	stack.Pop()
	if stack.Peek() != 1 {
		t.Fatal("Incorrect value.")
	}
}

func TestStackPush(t *testing.T) {
	stack := new(common.Stack)
	stack.Push(0)
	if stack.Peek() != 0 {
		t.Fatal("Incorrect value.")
	}
}

func TestStackLen(t *testing.T) {
	stack := new(common.Stack)
	for i := 0; i < 1024; i++ {
		stack.Push(i)
	}
	if stack.Len() != 1024 {
		t.Fatal("Incorrect value.")
	}
}
