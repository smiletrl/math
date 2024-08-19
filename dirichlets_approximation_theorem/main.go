package main

import (
	"log"
	"math"
)

func main() {
	ps := []float64{
		math.Sqrt(3),
		math.Cbrt(3),
		math.Pow(math.Pi, 2),
		math.Pow(math.E, 3),
	}
	n := 10

	for _, p := range ps {
		a, b := cal(n, p)
		v := equationVerify(a, b, n, p)
		if !v {
			panic("equation parameter is wrong")
		}
		log.Printf("Equation a: %d, b: %d, p: %.5f", a, b, p)
	}
}

func cal(n int, p float64) (chosenA, chosenB int) {
	min := float64(-1)
	for a := 1; a <= n; a++ {
		x := float64(a) * p
		f := math.Abs(x - math.Round(x))
		if min == -1 {
			min = f
			chosenA = a
			chosenB = int(math.Round(x))
		} else if min > f {
			min = f
			chosenA = a
			chosenB = int(math.Round(x))
		}
	}
	return
}

func equationVerify(chosenA, chosenB, n int, p float64) bool {
	if math.Abs(float64(chosenA)*p-float64(chosenB)) < float64(1/float64(n)) {
		return true
	}
	return false
}
