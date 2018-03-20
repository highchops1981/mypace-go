
//https://projecteuler.net/archives
package main

import (
	"fmt"
	"strconv"
)

// 4
type Calc interface {
	calc() int
}

type Multiplication struct {
	a int
	b int
}

func (m *Multiplication) calc() int {
	return m.a * m.b
}

type StringProcess interface {
	process() string
}

type Reverse struct {
	c string
}

func (s *Reverse) process() string {
	arr := []rune(s.c)
	for i,j := 0,len(arr)-1; i < j; i,j = i+1,j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	
	return string(arr)
}

func getMax() {

	max := 0
	palindromic := 0

	for i := 100; i < 1000; i++ {
	for j := 100; j < 1000; j++ {
		
		multiplication := Multiplication{i,j}
		c := multiplication
		palindromic = c.calc()
		
		reverse := Reverse{strconv.Itoa(palindromic)}
		s := reverse
		palindromicr,err := strconv.Atoi(s.process())
		if err == nil {
		}
		
		if  palindromic == palindromicr && max <= palindromic {
			max = palindromic
		}
	
	}}
	
	fmt.Println(max)
	
}

func main() {
	getMax()	
}


// 3
func main() {
	
	const max int64 = 600851475143
	var quotient int64 = 600851475143
	var remainder int64 = 0
	var denominator int64 = 2
	
	for {
	
		if checkPrimeNumber(denominator, max) == true {
		
			remainder = quotient % denominator
			
			if remainder == 0 {
			
				quotient = quotient / denominator
			
				if quotient == 1 {
				
					break
					
				}
			
			} else {
			
				denominator += 1
			
			}
			
		
		} else {
		
			denominator += 1
		
		}
		
	}
	
	fmt.Println(denominator)
	
	
}

func checkPrimeNumber(target int64, max int64) (bool) {

	isPrimeNumber := true
	var i int64 = 1
	
	for {
	
		if i > target || i > max {
		
			break
		}
		
		if i != 1 && i != target {
		
			if target % i == 0 {
			
				isPrimeNumber = false
				break
			}
		}
		
		i += 1
		
	}
	
	return isPrimeNumber

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

