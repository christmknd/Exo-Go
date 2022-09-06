package main

import (
	"fmt"
)

type Coors struct {
	x, y int
}

type ship struct {
	coors       Coors
	orientation int
	size        int
	hits        []Coors
}

func generation() {

}

func display() {
	var i, j int
	a := 0
	l := 10

	for j := 0; j < 10; j++ {
		a++
		fmt.Printf(" %d", a)
	}
	a = l
	fmt.Printf("\n")
	for i = 0; i < l; i++ {
		fmt.Printf("|")
		for j = 0; j < 10; j++ {
			fmt.Printf(" |")
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func main() {
	display()
}
