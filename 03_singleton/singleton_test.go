package main

import (
	"testing"
	"sync"
)

const parCount = 100

func TestSingleton(t *testing.T) {
	instance1 := GetInstance()
	instance2 := GetInstance()

	if instance1 != instance2 {
		t.Fatal("instance is not equal")
	}
}

func TestParallelSingleton(t *testing.T) {
	start := make(chan struct{})
	wg := sync.WaitGroup{}
  instances := [parCount]Singleton{}
	wg.Add(parCount)
	for i := 0; i < parCount; i++ {
		go func(index int) {
			<-start
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}
	close(start)
	wg.Wait()
	for i := 1; i < parCount; i++ {
		if(instances[i] != instances[i-1]) {
			t.Fatal("instance is not equal")
		}
	}
}
