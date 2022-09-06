package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	i := rand.Intn(10)
	j := rand.Intn(20)
	k := rand.Intn(100)
	fmt.Printf("|%d|,|%02d|, |%03d|\n", i, j, k)

	t := time.Now()
	month := t.Month()
	day := t.Day()
	hour := t.Hour()
	minutes := t.Minute()

	fmt.Printf("Nous sommes le %v %v , il est %v h %v\n  ", day, month, hour, minutes)

}
