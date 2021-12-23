package main 

import(
	"fmt"
)
type API interface{
	Say()
}

type hiAPI struct{
	hi string
}

func (a hiAPI) Say() {
	fmt.Println(a.hi)
}

type helloAPI struct {
	hello string
}

func (a helloAPI) Say() {
	fmt.Println(a.hello)
}

func GetAPI(i int) API {
	if i == 1 {
		return hiAPI{"hi"}
	} else if i == 2 {
		return helloAPI{"hello"}
	}
	return nil
}

func main() {
}