package main

import (
	"fmt"
	"sync"
)

type Singleton interface{
		foo()
}

type singleton struct{

}

func (s singleton) foo() {
	fmt.Println("singleton foo()")	
}

var(
	instance *singleton
	once			sync.Once
)

func GetInstance() Singleton {
		once.Do(func(){
			instance = &singleton{}
		})

		return instance
}

func main() {
		var s = GetInstance()
		s.foo()
}

