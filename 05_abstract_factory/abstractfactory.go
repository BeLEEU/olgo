package main

import (
	"fmt"
)

type Fruit interface {
	BecomeJuice()
}

type Vegetables interface {
	BecomeLunch()
}

type FoodFactory interface {
	CreateFruit() Fruit
	CreateVegetables() Vegetables
}

type Apple struct{}

func (*Apple) BecomeJuice() {
	fmt.Println("I will become Juice!")
}

type Oinion struct{}

func (*Oinion) BecomeLunch() {
	fmt.Println("I am your nightmare!")
}

type LunchFactory struct{}

func (*LunchFactory) CreateFruit() Fruit {
	return &Apple{}
}

func (*LunchFactory) CreateVegetables() Vegetables {
	return &Oinion{}
}

func main() {

}
