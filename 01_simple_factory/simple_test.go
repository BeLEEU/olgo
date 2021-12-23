package main

import(
	"testing"
)

func TestSimple(t *testing.T) {
	hi := GetAPI(1)
	hello := GetAPI(2)
	hi.Say()
	hello.Say()
}