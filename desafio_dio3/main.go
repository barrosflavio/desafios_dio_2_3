package main

import (
	"fmt"	
)

func main() {
	n := 1
	for n <= 100 {
		if n % 3 == 0 && n % 5 == 0 {
			fmt.Printf("PIN PAN\n")
			n++
		} else if n % 3 == 0 {
			fmt.Printf("PIN\n")
			n ++
		} else if n % 5 == 0 {
			fmt.Printf("PAN\n")
			n ++
		} else {
			fmt.Printf(" %d\n", n)
			n ++
		}
	}
}