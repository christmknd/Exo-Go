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

func (myship *ship) generateShip(coors Coors, orientation int, size int, hits Coors) {
	myship.coors = Coors{coors.x, coors.y}
	myship.orientation = orientation
	myship.size = size
}

func (myship *ship) displayShip(coors Coors, orientation int, size int) {
	fmt.Printf("Cordonnées : %d | %d \n ", myship.coors.x, myship.coors.y)
	fmt.Printf("Orientation : %d \n ", myship.orientation)
	fmt.Printf("size: %d \n", myship.size)
	//fmt.Printf("hits: %d \n", myship.hits)
}

func displayGrid() {
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
	displayGrid()
}
