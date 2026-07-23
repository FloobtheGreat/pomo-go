package main

import (
	"fmt"
	"time"
)

type Timer struct {
	work string
	rest string
}

var (
	short    = Timer{"15m", "3m"}
	standard = Timer{"25m", "5m"}
	long     = Timer{"50m", "10m"}
)

func statusUpdate() string { return "" }

func main() {
	fmt.Printf("%v\n%v\n%v\n", short, standard, long)
	c := time.Tick(5 * time.Second)
	for next := range c {
		fmt.Printf("%v %s\n", next, statusUpdate())
	}

}
