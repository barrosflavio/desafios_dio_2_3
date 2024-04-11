package main

import (
	"fmt"	
)

func main() {
	n := 1
	for n <= 100 {
		if n % 3 == 0 {
			div := n/3
			fmt.Printf("%d: %d/3=%d\n", n, n, div)
			n ++
		} else {
			n ++
		}
	}
}