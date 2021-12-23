package main

import (
	"testing"
)

func getObj(factory FoodFactory) {
	factory.CreateFruit().BecomeJuice()
	factory.CreateVegetables().BecomeLunch()
}

func TestGetLunch(t *testing.T) {
	
	factory := &LunchFactory{}
	getObj(factory)
}
