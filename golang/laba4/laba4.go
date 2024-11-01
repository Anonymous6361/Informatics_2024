package laba4

import (
	"fmt"
	"math"
)

func Colcul(a, b, x float64) float64 {
	y := (math.Log(math.Pow(b, 2)-math.Pow(x, 2)) / math.Log(a)) / (math.Pow(x*x-a*a, 1/3))
	return y
}
func TaskA(a, b, xMin, xMax, xDelta float64) []float64 {
	var result []float64
	for i := xMin; i < xMax; i += xDelta {
		result = append(result, Colcul(a, b, i))
	}
	return result
}

func TaskB(a, b float64, x []float64) []float64 {
	var result []float64
	for _, i := range x {
		result = append(result, Colcul(a, b, i))
	}
	return result
}

func CompleteLaba4() {
	var a float64 = 2.0
	var b float64 = 4.1
	var xMin float64 = 0.77
	var xMax float64 = 1.77
	var xDelta float64 = 0.2
	var x []float64 = []float64{1.24, 1.38, 2.38, 3.21, 0.68}

	var resultA []float64 = TaskA(a, b, xMin, xMax, xDelta)
	fmt.Println(resultA)
	var resultB []float64 = TaskB(a, b, x)
	fmt.Println(resultB)
}
