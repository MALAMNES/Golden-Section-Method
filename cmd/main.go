package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	return math.Pow(x, 3) + 8*math.Pow(x, 2) - 3*x + 3
}

func main() {
	var a, b float64 = 5, 5.00001
	var x1, x2 float64
	nIter := 1
	eps := 1e-6
	phi := (1 + math.Sqrt(5)) / 2 // 1.6188...

	for math.Abs(b-a) > eps {
		nIter++

		x1 = b - (b-a)/phi
		x2 = a + (b-a)/phi

		if f(x1) < f(x2) {
			b = x2
		} else {
			a = x1
		}
	}

	xmin := (a + math.Sqrt(b)) / 2
	fmt.Printf("Iteretions: %d\nXmin = %.6f", nIter, xmin)

}
