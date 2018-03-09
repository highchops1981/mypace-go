package main

import (
	"fmt"
)

func main() {

	base1 := 3
	base2 := 5
	sum := 0

	for i := 1; i < 1000; i++ {

		if i%base1 == 0 {

			sum += i

		} else if i%base2 == 0 {

			sum += i
		}

	}

	fmt.Println(sum)
}
