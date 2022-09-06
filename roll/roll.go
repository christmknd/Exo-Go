package main

import (
	"fmt"
	"math/rand"
)

func main() {
	i := rand.Intn(10)
	j := rand.Intn(20)
	k := rand.Intn(100)
	fmt.Printf("|%d|,|%02d|, |%03d|\n", i, j, k)

}
