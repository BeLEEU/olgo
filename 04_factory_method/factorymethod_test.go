package main

import (
	"testing"
)

func compute(factory OperatorFactory, x, y int) int {
	op := factory.Create()
	op.SetA(x)
	op.SetB(y)
	return op.Result()
}

func TestFactoryMethod(t *testing.T) {
	var (
		factory OperatorFactory
	)
	factory = PlusOperatorFactory{}
	if compute(factory, 2, 1) != 3 {
		t.Fatal("error with factory method pattern")
	}
	factory = MinusOperatorFactory{}
	if compute(factory, 2, 1) != 1 {
		t.Fatal("error with factory method pattern")
	}

}
