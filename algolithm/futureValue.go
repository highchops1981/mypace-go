package main

import "fmt"
import "math"

func main() {

	var pYen int = 1000000
	var pRate int = 5
	var pYaer int = 2
	var futureValue = calcFutureValue(pYen,pRate,pYaer)

	fmt.Printf("futureValue is %f.",futureValue)
	
}

func calcFutureValue(pYen int,pRate int,pYaer int) float64 {

	var futureValue = 0.0

	futureValue = float64(pYen) * math.Pow((1.0 + float64(pRate)/100.0),float64(pYaer))

	return futureValue

}
