package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	month := t.Month()
	day := t.Day()
	hour := t.Hour()
	minutes := t.Minute()

	fmt.Printf("Nous sommes le %v %v , il est %v h %v\n  ", day, month, hour, minutes)

}
