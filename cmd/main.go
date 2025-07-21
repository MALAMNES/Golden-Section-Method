package main

import (
	"fmt"
	"math"
)

func goldenSection(f func(float64) float64, a, b, eps float64) (int, float64) {

	phi := (1 + math.Sqrt(5)) / 2 // 1.6188...
	var xMin float64
	nIter := 0

	var recurse func(a, b float64)
	recurse = func(a, b float64) {
		nIter++

		if math.Abs(b-a) < eps {
			xMin = (a + math.Sqrt(b)) / 2

			return
		}

		x1 := b - (b-a)/phi
		x2 := a + (b-a)/phi

		if f(x1) < f(x2) {
			recurse(a, x2)
		} else {
			recurse(x1, b)
		}
	}

	recurse(a, b)

	return nIter, xMin
}

func main() {
	f := func(x float64) float64 { return math.Pow(x, 3) + 8*math.Pow(x, 2) - 3*x + 3 }

	var a, b float64 = 5, 5.00001
	eps := 1e-6

	nIter, xMin := goldenSection(f, a, b, eps)
	fmt.Printf("Iteretions: %d\nXmin = %.6f", nIter, xMin)

}
