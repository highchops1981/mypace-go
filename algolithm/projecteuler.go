package main

import (
	"fmt"
)

//1
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

//2
func main() {
	
	max := 4000000
	sum := 0
	kou := 0
	const n1 = 1
	const n2 = 2
	lastKou := 0
	lastlastKou := 0
	
	for kou <= max {
	
		if kou == 0 {
			
			kou = n1 + n2
			lastlastKou = n2
			lastKou = kou
		
		} else {
		
			kou = lastlastKou + lastKou
			lastlastKou = lastKou
			lastKou = kou
		
		}
		
		if kou % 2 == 0 {
		
			sum += kou
		}
	
	}
	
	sum += n2
	fmt.Println(sum)
		
}
